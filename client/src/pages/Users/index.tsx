import React, { useEffect, useMemo, useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import {
  DataGrid,
  GridColDef,
  GridValueGetterParams,
  GridToolbarContainer,
  GridSelectionModel,
  GridRowParams,
} from "@mui/x-data-grid";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { toJS } from "mobx";
import UserForm from "./UserForm";
import { roleByID } from "stores/UserStore";

export type DialogModelType = "add" | "change" | undefined;

type Props = {};

const Users = (props: Props) => {
  const { userStore } = useStore();

  const columns: GridColDef[] = [
    { field: "id", headerName: "ID", width: 30, hide: true },
    { field: "firstName", headerName: "First name", width: 130 },
    { field: "lastName", headerName: "Last name", width: 130 },
    { field: "age", headerName: "Age", width: 90 },
    {
      field: "role",
      headerName: "User Role",
      // description: "This column has a value getter and is not sortable.",
      sortable: false,
      width: 160,
      valueGetter: (params: GridValueGetterParams) =>
        roleByID(params.row.roleId),
    },
    { field: "email", headerName: "Email address", flex: 1, minWidth: 100 },
    {
      field: "office",
      headerName: "Office",
      width: 130,
      valueGetter: (params: GridValueGetterParams) =>
        userStore.officeByID(params.row.officeId)?.title,
    },
  ];

  useEffect(() => {
    if (userStore) {
      userStore.getUsers();
      userStore.getOffices();
    }

    return;
  }, []);

  const [selectionModel, setSelectionModel] = useState<GridSelectionModel>([]);

  const [dialogModel, setDialogModel] = useState<DialogModelType>();

  const handleClose = () => {
    if (!!dialogModel) setDialogModel(undefined);
  };

  return (
    <>
      <DataGrid
        rows={toJS(userStore.users)}
        columns={columns}
        autoHeight
        loading={userStore.status === "pending"}
        sx={{
          "&.MuiDataGrid-root .MuiDataGrid-columnHeader:focus-within, &.MuiDataGrid-root .MuiDataGrid-cell:focus-within":
            {
              outline: "none !important",
            },
          "& .is-active-user--false": {
            bgcolor: (theme) => theme.palette.error.light,
          },
        }}
        getRowClassName={(params) => `is-active-user--${params.row.active}`}
        selectionModel={selectionModel}
        onSelectionModelChange={(newModel, opt) => {
          setSelectionModel(newModel);
        }}
        components={{
          Toolbar: () => (
            <GridToolbarContainer
              sx={{ justifyContent: "space-between", gap: 2 }}
            >
              <Button
                color="info"
                variant="contained"
                onClick={() => {
                  setDialogModel("add");
                }}
              >
                Add user
              </Button>
              <Typography align="center" variant="h5">
                Users
              </Typography>
              <Box
                display="flex"
                gap={3}
                sx={{
                  overflow: "hidden",
                  textAlign: "center",
                  width: !selectionModel.length ? "0px" : undefined,
                  height: !selectionModel.length ? "0px" : undefined,
                }}
              >
                <Button
                  color="warning"
                  variant="contained"
                  onClick={() => {
                    setDialogModel("change");
                  }}
                >
                  Change role
                </Button>
                <Button
                  color={
                    userStore.userByID(selectionModel[0])?.active
                      ? "error"
                      : "success"
                  }
                  variant="contained"
                  onClick={() => userStore.switchActive(selectionModel[0])}
                >
                  {userStore.userByID(selectionModel[0])?.active
                    ? "Disable"
                    : "Enable"}
                </Button>
              </Box>
            </GridToolbarContainer>
            // BaseSelect // to visually update active status
          ),
        }}
      />
      <UserForm
        model={dialogModel}
        handleClose={handleClose}
        userId={useMemo(() => selectionModel[0], [selectionModel[0]])}
      />
    </>
  );
};

export default observer(Users);

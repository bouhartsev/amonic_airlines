import { useEffect, useMemo, useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import {
  DataGrid,
  GridColDef,
  GridValueGetterParams,
  GridToolbarContainer,
  GridSelectionModel,
  ValueOptions,
  GridValueFormatterParams,
} from "@mui/x-data-grid";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { toJS } from "mobx";
import UserForm from "./UserForm";
import { rolesOptions, roleByID } from "stores/UserStore";
import { tableBaseSX } from "utils/theme";

export type DialogModelType = "add" | "change" | undefined;

type Props = {};

const Users = (props: Props) => {
  const { userStore } = useStore();

  const columns: GridColDef[] = [
    { field: "id", headerName: "ID", width: 30, hide: true },
    { field: "firstName", headerName: "First name", width: 130 },
    { field: "lastName", headerName: "Last name", width: 130 },
    { field: "age", type: "number", headerName: "Age", width: 90 },
    {
      field: "roleId",
      headerName: "User Role",
      type: "singleSelect",
      valueOptions: () =>
        rolesOptions.map(
          (el) =>
            ({
              value: el.id,
              label: el.label,
            } as ValueOptions)
        ),
      sortable: false,
      width: 160,
      valueFormatter: (params: GridValueFormatterParams) =>
        roleByID(params.value),
      valueGetter: (params: GridValueGetterParams) =>
        params.row.roleId.toString(),
    },
    { field: "email", headerName: "Email address", flex: 1, minWidth: 100 },
    {
      field: "officeId",
      headerName: "Office",
      width: 130,
      type: "singleSelect",
      valueOptions: () =>
        userStore.offices.map(
          (el) =>
            ({
              value: el.id,
              label: el.title,
            } as ValueOptions)
        ),
      valueFormatter: (params: GridValueFormatterParams) =>
        userStore.officeByID(params.value)?.title,
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
        sx={tableBaseSX}
        getRowClassName={(params) => `row-status--${params.row.active}`}
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

import React, { useEffect, useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import {
  DataGrid,
  GridColDef,
  GridValueGetterParams,
  GridToolbarContainer,
  GridSelectionModel,
  GridRowParams
} from "@mui/x-data-grid";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import UserForm from "./UserForm";
import UserStore from "stores/UserStore";

type Props = {};

// function CustomToolbar() {
//   return (
//     <GridToolbarContainer sx={{ justifyContent: "space-between" }}>
//       {/* <GridToolbarColumnsButton /> */}
//       <Button>Add</Button>
//       <Typography align="center" variant="h5">
//         Users
//       </Typography>
//       <Box display={}>
//         <Button color="warning">Change role</Button>
//         <Button color={}>Disable</Button>
//       </Box>
//     </GridToolbarContainer>
//   );
// }

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
        userStore.roleByID(params.row.roleId),
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

  return (
    <>
      <DataGrid
        rows={userStore.users}
        columns={columns}
        autoHeight
        loading={
          userStore.status === "pending" || userStore.status === "initial"
        }
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
          setSelectionModel(newModel)}
        }
        components={{
          Toolbar: () => (
            <GridToolbarContainer
              sx={{ justifyContent: "space-between", gap: 2 }}
            >
              {/* <GridToolbarColumnsButton /> */}
              <Button color="info" variant="contained">
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
                <Button color="warning" variant="contained">
                  Change role
                </Button>
                <Button color={userStore.userByID(selectionModel[0])?.active ? "error" : "success"} variant="contained">
                {userStore.userByID(selectionModel[0])?.active ? "Disable" : "Enable"}
                </Button>
              </Box>
            </GridToolbarContainer>
          ),
        }}
      />
      <UserForm />
    </>
  );
};

export default observer(Users);

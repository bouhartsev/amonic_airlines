import { Box, Typography } from "@mui/material";
import { DataGrid, GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import { tableBaseSX } from "utils/theme";
import { toJS } from "mobx";
import React, { useEffect, useState } from "react";
import { useStore } from "stores";
import { options } from "yargs";

type Props = {};

const getStringOrNull = (params: GridValueGetterParams) =>
  params.row[params.field] || "--";
const getTimeOrNull = (params: GridValueGetterParams) =>
  !params.row[params.field]
    ? "--"
    : new Date(params.row[params.field]+" +0000").toLocaleString();

const columns: GridColDef[] = [
  {
    field: "loginTime",
    headerName: "Login Time",
    width: 170,
    valueGetter: getTimeOrNull,
  },
  {
    field: "logoutTime",
    headerName: "Logout time",
    width: 170,
    valueGetter: getTimeOrNull,
  },
  {
    field: "timeSpent",
    headerName: "Time spent on system",
    width: 100,
    valueGetter: getStringOrNull,
  },
  {
    field: "error",
    headerName: "Unsuccessful logout reason",
    description: "Why has it happened?",
    sortable: false,
    minWidth: 200,
    flex: 1,
    valueGetter: getStringOrNull,
  },
];

const Profile = (props: Props) => {
  const { userStore } = useStore();

  const [currentSessionTime, setCurrentSessionTime] = useState<string>("--:--:--");

  useEffect(() => {
    userStore.getProfile(userStore.userData.id);
    const watches = setInterval(
      () =>
        setCurrentSessionTime(
          new Date(
            Date.now() - userStore.profileData.iat * 1000
          ).toLocaleTimeString("eu", { timeZone: "UTC" }) // temp: change locale to normal for days
        ),
      1000
    );

    return () => {
      clearInterval(watches);
    };
  }, []);

  return (
    <>
      <Box sx={{ textAlign: "center" }}>
        <Typography component="h1" variant="h2">
          Hi, {userStore.userData.firstName}!
        </Typography>
        <Typography variant="overline" gutterBottom>
          Welcome to AMONIC Airlines
        </Typography>
      </Box>
      <Typography gutterBottom>
        Current session: {currentSessionTime}
      </Typography>
      <Typography gutterBottom>
        Number of crashes: {userStore.profileData.numberOfCrashes}
      </Typography>
      <DataGrid
        rows={toJS(userStore.profileData.userLogins || [])}
        columns={columns}
        autoHeight
        loading={userStore.status === "pending"}
        sx={tableBaseSX}
        getRowClassName={(params) =>
          `row-status--${!params.row.error && !!params.row.logoutTime}`
        }
        disableSelectionOnClick
      />
    </>
  );
};

export default Profile;

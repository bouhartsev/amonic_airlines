import { Box, Typography } from "@mui/material";
import {
  DataGrid,
  GridColDef,
  GridValueFormatterParams,
  GridValueGetterParams,
} from "@mui/x-data-grid";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { tableBaseSX } from "utils/theme";
import { toJS } from "mobx";
import React, { useEffect, useState } from "react";
import { useStore } from "stores";
import { observer } from "mobx-react-lite";
import UndetectedForm from "./UndetectedForm";

type Props = {};

const getStringOrNull = (params: GridValueGetterParams) =>
  params.row[params.field] || "--";
const getTimeOrNull = (params: GridValueGetterParams) =>
  !params.row[params.field]
    ? null
    : new Date(params.row[params.field] + " +0000");
const timeFormatter = (params: GridValueFormatterParams) =>
  !params.value ? "--" : params.value.toLocaleString();

const columns: GridColDef[] = [
  {
    field: "loginTime",
    type: "dateTime",
    headerName: "Login time",
    width: 170,
    valueGetter: getTimeOrNull,
    valueFormatter: timeFormatter,
  },
  {
    field: "logoutTime",
    type: "dateTime",
    headerName: "Logout time",
    width: 170,
    valueGetter: getTimeOrNull,
    valueFormatter: timeFormatter,
  },
  {
    field: "timeSpent",
    headerName: "Duration",
    description: "Time spent on system",
    width: 100,
    valueGetter: getStringOrNull,
  },
  {
    field: "error",
    headerName: "Crash reason",
    description: "Unsuccessful logout reason",
    sortable: false,
    minWidth: 200,
    flex: 1,
    valueGetter: getStringOrNull,
  },
];

const Profile = (props: Props) => {
  const { userStore } = useStore();

  const [undetected, setUndetected] = useState<boolean>(false);
  const [currentSessionTime, setCurrentSessionTime] =
    useState<string>("--:--:--");

  useEffect(() => {
    userStore.getProfile(userStore.userData.id).then(() => {
      if (!!userStore.profileData.LastLoginErrorDatetime) setUndetected(true);
    });
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
      <LocalizationProvider dateAdapter={AdapterDateFns}>
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
        <UndetectedForm
          open={undetected}
          handleClose={() => setUndetected(false)}
        />
      </LocalizationProvider>
    </>
  );
};

export default observer(Profile);

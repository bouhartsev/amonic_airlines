import React, { FC } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Button,
  Typography,
} from "@mui/material";
import { useStore } from "stores";

const PassengersList: FC = () => {
  const { bookingStore } = useStore();
  return (
    <>
      <Typography sx={{ my: 2 }} variant="subtitle1">
        Passengers list
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>First Name</TableCell>
              <TableCell>Last Name</TableCell>
              <TableCell>Birthdate</TableCell>
              <TableCell>Passport number</TableCell>
              <TableCell>Passport country</TableCell>
              <TableCell>Phone</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {bookingStore.passengers.map((row) => (
              <TableRow
                key={row.passportNumber}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.firstName}
                </TableCell>
                <TableCell align="left">{row.lastName}</TableCell>
                <TableCell align="left">
                  {row.birthdate.toLocaleDateString()}
                </TableCell>
                <TableCell align="left">{row.passportNumber}</TableCell>
                <TableCell align="left">
                  {bookingStore.countryByID(row.passportCountryId)?.name}
                </TableCell>
                <TableCell align="left">{row.phone}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Button
        variant="contained"
        color="error"
        disableElevation={true}
        sx={{ mt: 3 }}
      >
        Remove passenger
      </Button>
    </>
  );
};

export default PassengersList;

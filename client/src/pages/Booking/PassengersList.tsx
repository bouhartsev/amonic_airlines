import React, { FC } from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { makeStyles, withStyles } from '@mui/material';

function createData(
  name: string,
  calories: number,
  fat: number,
  carbs: number,
  protein: number
) {
  return { name, calories, fat, carbs, protein };
}

const rows = [
  createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
  createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
  createData('Eclair', 262, 16.0, 24, 6.0),
  createData('Cupcake', 305, 3.7, 67, 4.3),
  createData('Gingerbread', 356, 16.0, 49, 3.9),
];

const PassengersList: FC = () => {
  return (
    <Box>
      <Typography sx={{ mt: 5, mb: 2 }} variant="subtitle1">
        Passengers list
      </Typography>

      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>FirstName</TableCell>
              <TableCell>LastName</TableCell>
              <TableCell>Birthdate</TableCell>
              <TableCell>Passport Number</TableCell>
              <TableCell>Passport Country</TableCell>
              <TableCell>Phone</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.name}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="left">{row.calories}</TableCell>
                <TableCell align="left">{row.fat}</TableCell>
                <TableCell align="left">{row.carbs}</TableCell>
                <TableCell align="left">{row.protein}</TableCell>
                <TableCell align="left">{row.fat}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Button
        type="submit"
        variant="contained"
        className="btn-add"
        fullWidth={true}
        disableElevation={true}
        sx={{ mt: 3, width: '25%' }}
      >
        Remove passenger
      </Button>

      <Box
        sx={{
          mt: 3,
          display: 'flex',
          flexDirection: 'row',
          alignItems: 'center',
        }}
      >
        <Button variant="contained" color="secondary">
          Back to search for flights
        </Button>
        <Button sx={{ ml: 2 }} variant="contained" color="success">
          Confirm booking
        </Button>
      </Box>
    </Box>
  );
};

export default PassengersList;

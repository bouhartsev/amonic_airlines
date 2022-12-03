import React, { FC } from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Box from '@mui/material/Box';
import { styled } from '@mui/material/styles';
import { tableCellClasses } from '@mui/material/TableCell';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import NativeSelect from '@mui/material/NativeSelect';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableCellLight = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: theme.palette.action.hover,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  '&:nth-of-type(odd)': {
    backgroundColor: theme.palette.action.hover,
  },
  // hide last border
  '&:last-child td, &:last-child th': {
    border: 0,
  },
}));

const StyledTableRowQuestion = styled(TableRow)(({ theme }) => ({
  // '&:nth-of-type(odd)': {
  //   backgroundColor: theme.palette.action.hover,
  // },
  // hide last border
  '&:last-child td, &:last-child th': {
    border: 0,
  },
}));

function createData(
  question: String,
  answer: String,
  total: number,
  male: number,
  female: number,
  age1: number,
  age2: number,
  age3: number,
  age4: number,
  type1: number,
  type2: number,
  type3: number,
  dest1: number,
  dest2: number,
  dest3: number,
  dest4: number,
  dest5: number
) {
  return {
    question,
    answer,
    total,
    male,
    female,
    age1,
    age2,
    age3,
    age4,
    type1,
    type2,
    type3,
    dest1,
    dest2,
    dest3,
    dest4,
    dest5,
  };
}

const rows = [
  createData(
    'Please rate our aircraft flown on AMONIC Airlines',
    'Outstanding',
    583,
    44,
    47,
    40,
    40,
    42,
    58,
    48,
    22,
    15,
    43,
    50,
    49,
    39,
    47
  ),
  createData(
    '',
    'Very Good',
    504,
    41,
    38,
    38,
    44,
    40,
    35,
    39,
    20,
    14,
    34,
    39,
    36,
    46,
    40
  ),
  createData(
    '',
    'Good',
    1087,
    85,
    85,
    78,
    84,
    82,
    93,
    87,
    46,
    24,
    77,
    89,
    85,
    85,
    87
  ),
  createData('', 'Adequate', 123, 11, 8, 9, 8, 14, 5, 9, 5, 4, 15, 7, 11, 8, 9),
  createData(
    '',
    'Needs Improvement',
    19,
    2,
    1,
    3,
    1,
    1,
    1,
    1,
    2,
    1,
    1,
    1,
    1,
    2,
    1
  ),
  createData('', 'Poor', 140, 13, 9, 12, 9, 15, 6, 10, 7, 3, 16, 8, 12, 10, 10),
  createData('', "Don't know", 60, 2, 6, 10, 7, 3, 1, 3, 3, 2, 7, 4, 4, 5, 3),
];

const TableList: FC = () => {
  const [gender, setGender] = React.useState('');
  const [age, setAge] = React.useState('');

  const handleChangeAge = (event: SelectChangeEvent) => {
    setAge(event.target.value);
  };

  const handleChangeGender = (event: SelectChangeEvent) => {
    setGender(event.target.value);
  };

  return (
    <Box>
      <Box>
        <FormControl sx={{ minWidth: 120, mb: 5 }}>
          <InputLabel variant="standard">Time period</InputLabel>
          <NativeSelect defaultValue={'July 2017'}>
            <option value={'June 2017'}>June 2017</option>
            <option value={'July 2017'}>July 2017</option>
            <option value={'August 2017'}>August 2017</option>
          </NativeSelect>
        </FormControl>
      </Box>

      <Table sx={{ minWidth: 1200 }} aria-label="customized table">
        <TableHead sx={{ borderBottom: 2, borderColor: 'black' }}>
          <TableRow>
            <StyledTableCellLight
              align="center"
              colSpan={5}
            ></StyledTableCellLight>
            <StyledTableCellLight
              align="center"
              colSpan={1}
            ></StyledTableCellLight>
            <TableCell align="center" colSpan={2}>
              Gender
            </TableCell>
            <StyledTableCellLight align="center" colSpan={5}>
              Age
            </StyledTableCellLight>
            <TableCell align="center" colSpan={3}>
              Cabin Type
            </TableCell>
            <StyledTableCellLight align="center" colSpan={5}>
              Destination Airport
            </StyledTableCellLight>
          </TableRow>

          <TableRow>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={5}>
              Questions
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              Total
            </TableCell>
            <StyledTableCellLight sx={{ p: 1.5 }} align="center" colSpan={1}>
              Male
            </StyledTableCellLight>
            <StyledTableCellLight sx={{ p: 1.5 }} align="center" colSpan={1}>
              Female
            </StyledTableCellLight>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              18-24
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              25-39
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={2}>
              40-59
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              60+
            </TableCell>
            <StyledTableCellLight sx={{ p: 1.5 }} align="center" colSpan={1}>
              Economy
            </StyledTableCellLight>
            <StyledTableCellLight sx={{ p: 1.5 }} align="center" colSpan={1}>
              Business
            </StyledTableCellLight>
            <StyledTableCellLight sx={{ p: 1.5 }} align="center" colSpan={1}>
              First
            </StyledTableCellLight>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              AUH
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              BAH
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              DOH
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              RYU
            </TableCell>
            <TableCell sx={{ p: 1.5 }} align="center" colSpan={1}>
              CAI
            </TableCell>
          </TableRow>
        </TableHead>

        <TableBody>
          {rows.map((row) => (
            <>
              <Box fontWeight="fontWeightBold">
                <TableRow>{row.question}</TableRow>
              </Box>

              <StyledTableRow key={row.total}>
                <StyledTableCell colSpan={5} align="center">
                  {row.answer}
                </StyledTableCell>
                <StyledTableCell align="center">{row.total}</StyledTableCell>
                <StyledTableCell align="center">{row.male}</StyledTableCell>
                <StyledTableCell align="center">{row.female}</StyledTableCell>

                <StyledTableCell align="center">{row.age1}</StyledTableCell>
                <StyledTableCell align="center">{row.age2}</StyledTableCell>
                <StyledTableCell align="center" colSpan={2}>
                  {row.age3}
                </StyledTableCell>
                <StyledTableCell align="center">{row.age4}</StyledTableCell>

                <StyledTableCell align="center">{row.type1}</StyledTableCell>
                <StyledTableCell align="center">{row.type2}</StyledTableCell>
                <StyledTableCell align="center">{row.type3}</StyledTableCell>

                <StyledTableCell align="center">{row.dest1}</StyledTableCell>
                <StyledTableCell align="center">{row.dest2}</StyledTableCell>
                <StyledTableCell align="center">{row.dest3}</StyledTableCell>
                <StyledTableCell align="center">{row.dest4}</StyledTableCell>
                <StyledTableCell align="center">{row.dest5}</StyledTableCell>
              </StyledTableRow>
            </>
          ))}
        </TableBody>

        <TableBody>
          {rows.map((row) => (
            <>
              <Box fontWeight="fontWeightBold">
                <TableRow>{row.question}</TableRow>
              </Box>

              <StyledTableRow key={row.total}>
                <StyledTableCell colSpan={5} align="center">
                  {row.answer}
                </StyledTableCell>
                <StyledTableCell align="center">{row.total}</StyledTableCell>
                <StyledTableCell align="center">{row.male}</StyledTableCell>
                <StyledTableCell align="center">{row.female}</StyledTableCell>

                <StyledTableCell align="center">{row.age1}</StyledTableCell>
                <StyledTableCell align="center">{row.age2}</StyledTableCell>
                <StyledTableCell align="center" colSpan={2}>
                  {row.age3}
                </StyledTableCell>
                <StyledTableCell align="center">{row.age4}</StyledTableCell>

                <StyledTableCell align="center">{row.type1}</StyledTableCell>
                <StyledTableCell align="center">{row.type2}</StyledTableCell>
                <StyledTableCell align="center">{row.type3}</StyledTableCell>

                <StyledTableCell align="center">{row.dest1}</StyledTableCell>
                <StyledTableCell align="center">{row.dest2}</StyledTableCell>
                <StyledTableCell align="center">{row.dest3}</StyledTableCell>
                <StyledTableCell align="center">{row.dest4}</StyledTableCell>
                <StyledTableCell align="center">{row.dest5}</StyledTableCell>
              </StyledTableRow>
            </>
          ))}
        </TableBody>

        <TableBody>
          {rows.map((row) => (
            <>
              <Box fontWeight="fontWeightBold">
                <TableRow>{row.question}</TableRow>
              </Box>

              <StyledTableRow key={row.total}>
                <StyledTableCell colSpan={5} align="center">
                  {row.answer}
                </StyledTableCell>
                <StyledTableCell align="center">{row.total}</StyledTableCell>
                <StyledTableCell align="center">{row.male}</StyledTableCell>
                <StyledTableCell align="center">{row.female}</StyledTableCell>

                <StyledTableCell align="center">{row.age1}</StyledTableCell>
                <StyledTableCell align="center">{row.age2}</StyledTableCell>
                <StyledTableCell align="center" colSpan={2}>
                  {row.age3}
                </StyledTableCell>
                <StyledTableCell align="center">{row.age4}</StyledTableCell>

                <StyledTableCell align="center">{row.type1}</StyledTableCell>
                <StyledTableCell align="center">{row.type2}</StyledTableCell>
                <StyledTableCell align="center">{row.type3}</StyledTableCell>

                <StyledTableCell align="center">{row.dest1}</StyledTableCell>
                <StyledTableCell align="center">{row.dest2}</StyledTableCell>
                <StyledTableCell align="center">{row.dest3}</StyledTableCell>
                <StyledTableCell align="center">{row.dest4}</StyledTableCell>
                <StyledTableCell align="center">{row.dest5}</StyledTableCell>
              </StyledTableRow>
            </>
          ))}
        </TableBody>

        <TableBody>
          {rows.map((row) => (
            <>
              <Box fontWeight="fontWeightBold">
                <TableRow>{row.question}</TableRow>
              </Box>

              <StyledTableRow key={row.total}>
                <StyledTableCell colSpan={5} align="center">
                  {row.answer}
                </StyledTableCell>
                <StyledTableCell align="center">{row.total}</StyledTableCell>
                <StyledTableCell align="center">{row.male}</StyledTableCell>
                <StyledTableCell align="center">{row.female}</StyledTableCell>

                <StyledTableCell align="center">{row.age1}</StyledTableCell>
                <StyledTableCell align="center">{row.age2}</StyledTableCell>
                <StyledTableCell align="center" colSpan={2}>
                  {row.age3}
                </StyledTableCell>
                <StyledTableCell align="center">{row.age4}</StyledTableCell>

                <StyledTableCell align="center">{row.type1}</StyledTableCell>
                <StyledTableCell align="center">{row.type2}</StyledTableCell>
                <StyledTableCell align="center">{row.type3}</StyledTableCell>

                <StyledTableCell align="center">{row.dest1}</StyledTableCell>
                <StyledTableCell align="center">{row.dest2}</StyledTableCell>
                <StyledTableCell align="center">{row.dest3}</StyledTableCell>
                <StyledTableCell align="center">{row.dest4}</StyledTableCell>
                <StyledTableCell align="center">{row.dest5}</StyledTableCell>
              </StyledTableRow>
            </>
          ))}
        </TableBody>
      </Table>

      <Box sx={{ display: 'flex',width: '60%', justifyContent: 'space-between', mt: 5, mx: 'auto'  }}>

        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <FormControlLabel
            value="Gender"
            control={<Checkbox />}
            label="Gender"
          />

          <FormControl variant="standard" sx={{ m: 1, minWidth: 200 }}>
            <InputLabel id="gender">
              All genders
            </InputLabel>
            <Select
              labelId="gender"
              id="demo-simple-select-standard"
              value={gender}
              onChange={handleChangeGender}
              label="Gender"
            >
              <MenuItem value="">
                <em>None</em>
              </MenuItem>
              <MenuItem value={10}>Ten</MenuItem>
              <MenuItem value={20}>Twenty</MenuItem>
              <MenuItem value={30}>Thirty</MenuItem>
            </Select>
          </FormControl>
        </Box>

        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <FormControlLabel
            value="Age"
            control={<Checkbox />}
            label="Age"
          />

          <FormControl variant="standard" sx={{ m: 1, minWidth: 200 }}>
            <InputLabel id="age">
              All ages
            </InputLabel>
            <Select
              labelId="age"
              id="demo-simple-select-standard"
              value={age}
              onChange={handleChangeAge}
              label="Age"
            >
              <MenuItem value="">
                <em>None</em>
              </MenuItem>
              <MenuItem value={10}>Ten</MenuItem>
              <MenuItem value={20}>Twenty</MenuItem>
              <MenuItem value={30}>Thirty</MenuItem>
            </Select>
          </FormControl>
        </Box>

      </Box>
    </Box>
  );
};

export default TableList;

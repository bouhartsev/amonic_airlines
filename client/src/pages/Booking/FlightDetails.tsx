import React, { FC } from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';

const FlightDetails: FC = () => {
  return (
    <Box display="flex" flexDirection="column" flexWrap="wrap">
      <Box>
        <Typography variant="subtitle1">Outbound flight details</Typography>
        <Box
          display="flex"
          flexDirection="row"
          p="1"
          m="1"
          justifyContent="space-between"
        >
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">From:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              CAI
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">To:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              AUH
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Cabin Type:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              Economy
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Date:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              11/10/2017
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Flight number:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              1908
            </Box>
          </Box>
        </Box>
      </Box>

      <Box>
        <Typography variant="subtitle1">Return flight details</Typography>
        <Box
          display="flex"
          flexDirection="row"
          p="1"
          m="1"
          justifyContent="space-between"
        >
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">From:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              AUH
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">To:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              CAI
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Cabin Type:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              Economy
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Date:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              11/15/2017
            </Box>
          </Box>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="subtitle1">Flight number:</Typography>
            <Box fontWeight="fontWeightBold" m={1}>
              1907
            </Box>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

export default FlightDetails;
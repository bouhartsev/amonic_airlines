import React, { FC } from 'react';
import Box from '@mui/material/Box';
import './FlightDetails.css';

export const FlightDetails: FC = () => {
  return (
    <div className="flight-details">
      <div className="outbound">
        <p>Outbound flight details</p>
        <Box
          display="flex"
          flexDirection="row"
          p="1"
          m="1"
          justifyContent="space-between"
        >
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>From: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              CAI
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>To: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              AUH
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Cabin Type: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              Economy
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Date: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              11/10/2017
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Flight number: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              1908
            </Box>
          </div>
        </Box>
      </div>

      <div className="return">
        <p>Return flight details</p>
        <Box
          display="flex"
          flexDirection="row"
          p="1"
          m="1"
          justifyContent="space-between"
        >
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>From: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              AUH
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>To: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              CAI
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Cabin Type: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              Economy
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Date: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              11/15/2017
            </Box>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <p>Flight number: </p>
            <Box fontWeight="fontWeightBold" m={1}>
              1907
            </Box>
          </div>
        </Box>
      </div>
    </div>
  );
};

import React, { FC } from 'react';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import TableList from './TableList';

const Surveys: FC = () => {

  return (
    <Container maxWidth="xl">
      <Typography variant="h4" sx={{ textAlign: 'center', marginBottom: 5 }}>
        Total Results
      </Typography>

      <TableList />
    </Container>
  );
};

export default Surveys;

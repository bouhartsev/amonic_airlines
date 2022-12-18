import React, { useState } from "react";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import TableList from "./TableList";
import { Box, ToggleButton, ToggleButtonGroup } from "@mui/material";

const Surveys = () => {
  const [view, setView] = useState("brief");
  const handleChange = (
    event: React.MouseEvent<HTMLElement>,
    newView: string
  ) => {
    if (newView !== null) setView(newView);
  };

  return (
    <>
      <Box sx={{ textAlign: "center", mb: 3 }}>
        <Typography variant="h4">View Surveys</Typography>
        <ToggleButtonGroup
          color="primary"
          value={view}
          exclusive
          onChange={handleChange}
          aria-label="View"
          sx={{mt:2}}
        >
          <ToggleButton value="brief">Results Summary</ToggleButton>
          <ToggleButton value="detailed">Detailed Results</ToggleButton>
        </ToggleButtonGroup>
      </Box>

      <TableList />
    </>
  );
};

export default Surveys;

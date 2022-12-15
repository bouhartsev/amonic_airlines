import React, { useState } from "react";
import SearchFlights from "./SearchFlights";
import BookingConf from "./BookingConf";
import BillingConf from "./BillingConf";

import { Stepper, Step, StepButton, Box, Button } from "@mui/material";

const steps = [
  { title: "Search for flights", component: <SearchFlights /> },
  { title: "Booking confirmation", component: <BookingConf /> },
  { title: "Billing confirmation", component: <BillingConf /> },
];

const Booking = () => {
  const [activeStep, setActiveStep] = useState(0);

  const handleNext = () => {
    if (activeStep < steps.length - 1) setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  return (
    <>
      <Stepper activeStep={activeStep} sx={{ mb: 3 }}>
        {steps.map((step, index) => {
          const stepProps: { completed?: boolean } = {};
          return (
            <Step key={step.title} {...stepProps}>
              <StepButton onClick={() => setActiveStep(index)}>
                {step.title}
              </StepButton>
            </Step>
          );
        })}
      </Stepper>

      {steps[activeStep].component}

      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
          mt: 2,
        }}
      >
        <Button
          disabled={activeStep === 0}
          onClick={handleBack}
          variant="contained"
          color="secondary"
        >
          Back
          {activeStep > 0 && " to " + steps[activeStep - 1].title.toLowerCase()}
        </Button>
        <Button onClick={handleNext} variant="contained" color="success">
          {activeStep === steps.length - 1
            ? "Issue tickets"
            : "To " + steps[activeStep + 1].title.toLowerCase()}
        </Button>
      </Box>
    </>
  );
};

export default Booking;

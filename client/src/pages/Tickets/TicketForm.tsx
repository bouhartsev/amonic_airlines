import React, { useEffect } from "react";
import {
  Box,
  Button,
  Dialog,
  DialogContent,
  DialogTitle,
  IconButton,
  Typography,
} from "@mui/material";
import {
  useForm,
  FormContainer,
  TextFieldElement,
  AutocompleteElement,
  DatePickerElement,
  DateTimePickerElement,
} from "react-hook-form-mui";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { Close as CloseIcon } from "@mui/icons-material";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { FlightType } from "stores/FlightStore";
import { LoadingButton } from "@mui/lab";

type Props = {
  open: boolean;
  flightId: FlightType["id"];
  handleClose: VoidFunction;
};

const ScheduleForm = (props: Props) => {
  const { flightStore } = useStore();
  const formContext = useForm<FlightType>();

  useEffect(() => {
    const currentFlight = !!props.flightId
      ? flightStore.scheduleByID(props.flightId)
      : {};

    formContext.reset({ ...currentFlight });

    return () => {};
  }, [flightStore, formContext, props.flightId]);

  const handleClose = () => {
    flightStore.status = "initial";
    flightStore.error = "";
    formContext.reset({});
    props.handleClose();
  };
  const thenClose = (act: Promise<any>) => {
    return act.then((res) => {
      if (flightStore.status === "success") handleClose();
    });
  };
  const handleSubmit = (data: FlightType) => {
    return thenClose(flightStore.updateSchedule(data));
  };

  return (
    <>
      <Dialog open={!!props.open} onClose={handleClose}>
        <DialogTitle>
          Edit schedule
          <IconButton
            aria-label="close"
            onClick={handleClose}
            sx={{
              position: "absolute",
              right: 8,
              top: 8,
              color: (theme) => theme.palette.grey[500],
            }}
          >
            <CloseIcon />
          </IconButton>
        </DialogTitle>

        <DialogContent>
          <LocalizationProvider dateAdapter={AdapterDateFns}>
            <FormContainer formContext={formContext} onSuccess={handleSubmit}>
              <AutocompleteElement
                required
                // matchId
                name="from"
                label="From"
                textFieldProps={{ margin: "normal" }}
                autocompleteProps={{
                  getOptionLabel: (opt: any) => opt.name,
                  disabled: true,
                }}
                options={flightStore.airports}
              />
              <AutocompleteElement
                required
                // matchId
                name="to"
                label="To"
                textFieldProps={{ margin: "normal" }}
                autocompleteProps={{
                  getOptionLabel: (opt: any) => opt.name,
                  disabled: true,
                }}
                options={flightStore.airports}
              />
              <TextFieldElement
                required
                disabled
                name="aircraft"
                label="Aircraft"
                margin="normal"
              />
              <Box>
                {/* there are issues with date enter */}
                <DatePickerElement
                  required
                  name="date"
                  label="Date"
                  // inputProps={{ fullWidth: true, margin: "normal" }}
                />
                <DateTimePickerElement
                  required
                  name="time"
                  label="Time"
                  isDate={false}
                  // fullWidth
                />
                <TextFieldElement
                  required
                  name="economyPrice"
                  label="Economy Price"
                  margin="normal"
                  type="number"
                />
              </Box>
              <Typography color="error">
                {flightStore.status === "error" && flightStore.error}
              </Typography>
              <Box sx={{ mt: 2, display: "flex", gap: 1 }}>
                <LoadingButton
                  type="submit"
                  loading={flightStore.status === "pending"}
                  variant="contained"
                  fullWidth
                >
                  Save
                </LoadingButton>
                <Button
                  type="reset"
                  onClick={handleClose}
                  variant="contained"
                  color="warning"
                >
                  Cancel
                </Button>
              </Box>
            </FormContainer>
          </LocalizationProvider>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default observer(ScheduleForm);

import React, { useEffect } from "react";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
  Paper,
  dialogClasses,
  Typography,
  LinearProgress,
} from "@mui/material";
import {
  useForm,
  FormContainer,
  TextFieldElement,
  AutocompleteElement,
  RadioButtonGroup,
  DatePickerElement,
  DateTimePickerElement,
} from "react-hook-form-mui";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { Close as CloseIcon } from "@mui/icons-material";
import { DialogModelType } from "./";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { FlightType } from "stores/FlightStore";
import { LoadingButton } from "@mui/lab";
import { uncamelize } from "utils/converters";

type Props = {
  open: boolean;
  handleClose: VoidFunction;
};

const ImportForm = (props: Props) => {
  const { flightStore } = useStore();

  const handleClose = () => {
    flightStore.status = "initial";
    flightStore.error = "";
    flightStore.upload.progress = 0;
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
      <Dialog open={!!props.open} onClose={handleClose} fullWidth>
        <DialogTitle>
          Import changes
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
          <Button variant="contained" component="label">
            Upload
            <input
              hidden
              accept=".csv"
              // multiple
              type="file"
              onChange={(e) => flightStore.uploadSchedules(e.target.files)}
            />
          </Button>
          <LinearProgress
            variant="determinate"
            value={flightStore.upload.progress}
            sx={{my: 2}}
          />
          <Box display={flightStore.status === "success" ? "block" : "none"}>
            {Object.keys(flightStore.upload.results).map((key) => (
              <Typography>
                {uncamelize(key)}: {flightStore.upload.results[key]}
              </Typography>
            ))}
          </Box>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default observer(ImportForm);

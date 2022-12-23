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
  CheckboxButtonGroup,
  TextFieldElement,
} from "react-hook-form-mui";
import { Close as CloseIcon } from "@mui/icons-material";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { ReportType } from "stores/UserStore";
import { LoadingButton } from "@mui/lab";
import { camelize } from "utils/converters";

const radioCrash = ["Software Crash", "System Crash"].map((el) => ({
  id: camelize(el),
  label: el,
}));

type Props = {
  open: boolean;
  handleClose: VoidFunction;
};

const UndetectedForm = (props: Props) => {
  const { userStore } = useStore();
  const formContext = useForm<ReportType>();

  const handleClose = () => {
    userStore.status = "initial";
    userStore.error = "";
    formContext.reset({});
    props.handleClose();
  };
  const thenClose = (act: Promise<any>) => {
    return act.then((res) => {
      if (userStore.status === "success") handleClose();
    });
  };
  const handleSubmit = (data: ReportType) => thenClose(userStore.report(data));

  return (
    <>
      <Dialog open={!!props.open} onClose={handleClose}>
        <DialogTitle>
          Report about problem
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
          <FormContainer formContext={formContext} onSuccess={handleSubmit}>
            <Typography color="info">
              No logout detected for your last login on{" "}
              {userStore.profileData.LastLoginErrorDatetime}
            </Typography>
            <TextFieldElement
              required
              label="Logout reason"
              name="reason"
              multiline
            />
            <CheckboxButtonGroup
              label="What was it?"
              name="radioCrash"
              options={radioCrash}
            />
            <Typography color="error">
              {userStore.status === "error" && userStore.error}
            </Typography>
            <Box sx={{ mt: 2, display: "flex", gap: 1 }}>
              <LoadingButton
                type="submit"
                loading={userStore.status === "pending"}
                variant="contained"
                fullWidth
              >
                Report
              </LoadingButton>
              <Button
                type="reset"
                onClick={handleClose}
                variant="contained"
                color="warning"
              >
                Later
              </Button>
            </Box>
          </FormContainer>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default observer(UndetectedForm);

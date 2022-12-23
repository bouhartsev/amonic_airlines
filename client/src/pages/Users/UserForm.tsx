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
} from "@mui/material";
import {
  useForm,
  FormContainer,
  PasswordElement,
  TextFieldElement,
  AutocompleteElement,
  RadioButtonGroup,
  DatePickerElement,
} from "react-hook-form-mui";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { Close as CloseIcon } from "@mui/icons-material";
import { DialogModelType } from "./";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { UserType, rolesOptions } from "stores/UserStore";
import { LoadingButton } from "@mui/lab";

type Props = {
  model: DialogModelType;
  userId?: UserType["id"];
  handleClose: VoidFunction;
};

const UserForm = (props: Props) => {
  const { userStore } = useStore();
  const formContext = useForm<UserType>();

  useEffect(() => {
    const currentUser =
      props.model === "change" && !!props.userId
        ? userStore.userByID(props.userId)
        : { roleId: rolesOptions[0].id };

    formContext.reset({
      ...currentUser,
      roleId: currentUser?.roleId.toString(),
    });

    return () => {};
  }, [userStore, formContext, props.userId, props.model]);

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
  const handleSubmit = (data: UserType) => {
    return thenClose(
      props.model === "change" && !!props.userId
        ? userStore.updateUser(data)
        : userStore.addUser(data)
    );
  };

  return (
    <>
      <Dialog open={!!props.model} onClose={handleClose}>
        <DialogTitle>
          {props.model === "change" ? "Change role" : "Add user"}
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
              <TextFieldElement
                required
                disabled={props.model !== "add"}
                type="email"
                name="email"
                label="Email"
                fullWidth
                margin="normal"
              />
              <TextFieldElement
                required
                disabled={props.model !== "add"}
                type="text"
                name="firstName"
                label="First name"
                fullWidth
                margin="normal"
              />
              <TextFieldElement
                required
                disabled={props.model !== "add"}
                type="text"
                name="lastName"
                label="Last name"
                fullWidth
                margin="normal"
              />
              <AutocompleteElement
                required
                matchId
                name="officeId"
                label="Office"
                textFieldProps={{ margin: "normal" }}
                autocompleteProps={{
                  getOptionLabel: (opt: any) => opt.title,
                  disabled: props.model !== "add",
                }}
                options={userStore.offices}
              />
              <Box className={props.model !== "change" ? "" : "hidden"}>
                {/* there are issues with date enter */}
                <DatePickerElement
                  required={props.model === "add"}
                  name="birthdate"
                  label="Birthday"
                  inputProps={{ fullWidth: true, margin: "normal" }}
                />
                <PasswordElement
                  required={props.model === "add"}
                  disabled={props.model !== "add"}
                  name="password"
                  label="Password"
                  fullWidth
                  margin="normal"
                />
              </Box>
              <Box className={props.model === "change" ? "" : "hidden"}>
                <RadioButtonGroup
                  required
                  name="roleId"
                  label="Role"
                  type="string"
                  options={rolesOptions}
                />
              </Box>
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

export default observer(UserForm);

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
import { userType, roles } from "stores/UserStore";
import styles from "./index.module.css";

const rolesObj = roles.map((el, ind) => ({ id: ind + 1, label: el })).reverse();

type Props = {
  model: DialogModelType;
  userId?: userType["id"];
  handleClose: VoidFunction;
};

// const FormContainer = (props) => {

// }

const UserForm = (props: Props) => {
  const { userStore } = useStore();
  const formContext = useForm<userType>();

  useEffect(() => {
    const currentUser =
      props.model === "change" && !!props.userId
        ? userStore.userByID(props.userId)
        : { roleId: rolesObj[0].id };

    console.log({ ...currentUser });
    formContext.reset(currentUser);

    return () => {};
  }, [userStore, formContext, props.userId, props.model]);

  const handleSubmit = (data: userType) => {
    // userStore
    console.log(data);
  };
  const handleClose = () => {
    formContext.reset({});
    props.handleClose();
  };

  return (
    <>
      <Dialog open={!!props.model} onClose={handleClose}>
        {/*  PaperComponent={({children})=>{return <FormContainer formContext={formContext} onSuccess={handleSubmit}><Paper>{children}</Paper></FormContainer>}} */}
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
              <Box>
                <DatePickerElement
                  required
                  name="birthdate"
                  label="Birthday"
                  inputProps={{ fullWidth: true, margin: "normal" }}
                />
                <PasswordElement
                  required
                  disabled={props.model !== "add"}
                  name="password"
                  label="Password"
                  fullWidth
                  margin="normal"
                />
              </Box>
              <Box className={props.model === "change" ? "" : styles.hidden}>
                <RadioButtonGroup
                  required
                  name="roleId"
                  label="Role"
                  options={rolesObj}
                />
              </Box>
              <Box sx={{ mt: 2, display: "flex", gap: 1 }}>
                <Button type="submit" variant="contained" fullWidth>
                  Save
                </Button>
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

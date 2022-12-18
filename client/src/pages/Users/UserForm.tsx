import React from "react";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
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

const rolesObj = roles
  .map((el, ind) => ({ id: (ind + 1).toString(), label: el }))
  .reverse();

type Props = {
  model: DialogModelType;
  userId?: userType["id"];
  handleClose: VoidFunction;
};

const UserForm = (props: Props) => {
  const { userStore } = useStore();
  const currentUser =
    props.model === "change" && !!props.userId
      ? userStore.userByID(props.userId)
      : { roleId: rolesObj[0].id };
  const formContext = useForm<userType>({
    defaultValues: currentUser, // doesn't work yet
  });


  const handleSubmit = () => {
    // userStore
  };
  const handleClose = () => {
    formContext.reset({});
    props.handleClose();
  };

  return (
    <>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <FormContainer formContext={formContext} onSuccess={handleSubmit}>
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
                name="officeId"
                label="Office"
                textFieldProps={{ margin: "normal" }}
                autocompleteProps={{
                  getOptionLabel: (opt: any) => opt.title,
                  disabled: props.model !== "add",
                }}
                options={userStore.offices}
              />
              <Box className={props.model === "add" ? "" : styles.hidden}>
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
            </DialogContent>
            <DialogActions sx={{ mx: 2 }}>
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
              {/* <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleClose}>Subscribe</Button> */}
            </DialogActions>
          </Dialog>
        </FormContainer>
      </LocalizationProvider>
    </>
  );
};

export default observer(UserForm);

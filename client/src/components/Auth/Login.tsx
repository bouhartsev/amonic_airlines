import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { Navigate, Link as RouterLink } from "react-router-dom";
import { Box, Avatar, Typography, Button } from "@mui/material";
import { LoadingButton } from "@mui/lab";
import { LockOutlined } from "@mui/icons-material";
import {
  useForm,
  FormContainer,
  TextFieldElement,
  PasswordElement,
} from "react-hook-form-mui";

type Fields = { email: string; password: string };

const Login = () => {
  const { userStore } = useStore();

  const formContext = useForm<Fields>();

  if (!!userStore.isLogged) {
    if (userStore.userData?.role === "administrator")
      return <Navigate to="/users" />;
    return <Navigate to="/profile" />;
  }

  // const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
  //   event.preventDefault();
  //   const data = new FormData(event.currentTarget);

  //   userStore.login(String(data.get("email")), String(data.get("password")));
  // };
  const handleSubmit = (data: Fields) => {
    userStore.login(data.email, data.password);
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        m: "auto",
        width: 400,
      }}
    >
      <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
        <LockOutlined />
      </Avatar>
      <Typography component="h1" variant="h5">
        Login
      </Typography>
      <FormContainer formContext={formContext} onSuccess={handleSubmit}>
        <TextFieldElement
          required
          type="email"
          name="email"
          label="Email"
          fullWidth
          margin="normal"
        />
        <PasswordElement
          required
          name="password"
          label="Password"
          fullWidth
          margin="normal"
        />
        <Typography color="error">
          {userStore.status === "error" && userStore.error}
        </Typography>
        <Box sx={{ my: 3, display: "flex", gap: 1 }}>
          <LoadingButton
            type="submit"
            fullWidth
            loading={userStore.status === "pending"}
            disabled={userStore.status === "forbidden"}
            variant="contained"
          >
            Sign In
          </LoadingButton>
          <Button component={RouterLink} to="/" variant="contained" color="secondary">Exit</Button>
        </Box>
      </FormContainer>
    </Box>
  );
};

export default observer(Login);

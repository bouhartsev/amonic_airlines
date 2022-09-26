import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { Navigate } from "react-router-dom";
import { Box, Avatar, Typography, Button } from "@mui/material";
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
    if (userStore.userData.role === "Administrator")
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
      }}
      maxWidth="sm"
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
          {userStore.error}
        </Typography>
        <Button
          type="submit"
          // fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Sign In
        </Button>
      </FormContainer>
    </Box>
  );
};

export default observer(Login);

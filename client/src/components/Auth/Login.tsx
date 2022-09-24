import { useStore } from "stores";
import { Navigate } from "react-router-dom";

const Login = () => {
  const { userStore } = useStore();
  if (!!userStore.isLogged) {
    if (userStore.userData.role === "Administrator")
      return <Navigate to="/users" />;
    return <Navigate to="/profile" />;
  }
  return <div>Login</div>;
};

export default Login;

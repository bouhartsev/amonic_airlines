import { useEffect } from "react";
import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import { Navigate } from "react-router-dom";

const Logout = () => {
  const { userStore } = useStore();

  useEffect(() => {
    if (userStore.isLogged && userStore.status!=="pending")
      userStore.logout();

    return () => {};
  }, [userStore]);

  if (!!userStore.isLogged) return <div>Logging out...</div>;
  return <Navigate to="/" />;
};

export default observer(Logout);

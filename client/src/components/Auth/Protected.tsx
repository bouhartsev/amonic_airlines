import { useStore } from "stores";
import { Navigate, Outlet } from "react-router-dom";
import { observer } from "mobx-react-lite";

type Props = {
  role?: string,
}

const Protected = (props: Props) => {
  const { userStore } = useStore();
  if (!userStore.isLogged) return <Navigate to="/login" />;
  else if (!!props.role && props.role!==userStore.userData?.role) throw new Error("403");
  return <Outlet />;
};

export default observer(Protected);

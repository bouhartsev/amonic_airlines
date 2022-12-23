import React, { Component, ReactNode } from "react";
import ErrorPage from "pages/Error";
import { Alert, Snackbar } from "@mui/material";
interface Props {
  children: ReactNode;
}
interface State {
  error: string | null;
  showSmall?: boolean;
  showBig?: boolean;
}

class ErrorBoundary extends Component<Props, State> {
  public state: State = {
    error: null,
    showSmall: false,
    showBig: false,
  };
  public static getDerivedStateFromError(error: Error): State {
    // Update state so the next render will show the fallback UI.
    return { error: error.message, showBig: true };
  }
  // public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
  //   console.error("Uncaught error:", error, errorInfo);
  //   this.setState({ error: error, errorInfo: errorInfo });
  // }
  public componentDidMount(): void {
    const showError = (error: any) => {
      // add all axios handler
      // add 401 handler
      this.setState({ error: String(error), showSmall: true });
    };
    window.onerror = showError;
    window.onunhandledrejection = (event) => showError(event.reason);
  }
  // public componentDidUpdate(
  //   prevProps: Readonly<Props>,
  //   prevState: Readonly<State>,
  //   snapshot?: any
  // ): void {
  //   if (window.location.pathname != this.path) {
  //     this.setState({ showSmall: false });
  //     this.path = window.location.pathname;
  //   }
  // }
  public handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    // if (reason === "clickaway") {
    //   return;
    // }

    this.setState({ showSmall: false });
  };
  public render() {
    return (
      <>
        {!!this.state.showBig ? (
          <ErrorPage code={String(this.state.error)} />
        ) : (
          this.props.children // what with rerendering, remounting? May be needed useMemo or something...
        )}
        <Snackbar
          open={!!this.state.showSmall}
          autoHideDuration={6000}
          onClose={this.handleClose}
        >
          <Alert
            onClose={this.handleClose}
            severity="error"
            sx={{ width: "100%" }}
            elevation={6}
            variant="filled"
          >
            {this.state.error}
          </Alert>
        </Snackbar>
      </>
    );
  }
}
export default ErrorBoundary;

import { createTheme } from "@mui/material/styles";

// all colors from brandbook
// some colors (not mentioned) are inherited from lib - https://mui.com/material-ui/customization/palette/#default-values
// * "not bb" - not from brandbook

const theme = createTheme({
    typography: {
        fontFamily: [
            '"Tex Gyre Adventor"',
            '"Twentieth Century"',
            '"Verdana"',
            'sans-serif',
        ].join(','),
    },
    palette: {
        primary: {
            main: "#196AA6",
            dark: "#064B66",
        },
        secondary: {
            light: "#EDD688",
            main: "#F79420",
            dark: "#C2912E",
            contrastText: "#FFFFFF"
        },
        warning: {
            light: "#FFFACB",
            main: "#FAC826",
            dark: "#e6b207", // not bb
            contrastText: "#000000"
        },
        info: {
            main: "#00A088",
            dark: "#132B4F",
        },
        success: {
            main: "#2e7d32", // not bb
            dark: "#004F4C",
        },
    },
    components: {
        MuiCssBaseline: {
            styleOverrides: require("assets/fonts.css"),
        },
    },
});

export const tableBaseSX = {
    "&.MuiDataGrid-root .MuiDataGrid-columnHeader:focus-within, &.MuiDataGrid-root .MuiDataGrid-cell:focus-within":
    {
        outline: "none !important",
    },
    "& .row-status--false": {
        color: "error.contrastText",
    },
    "& .row-status--false::before": {
        bgcolor: "error.light",
    },
}

export default theme;
import axios from "axios";

export const api = axios.create({
    baseURL: process.env.REACT_APP_BASE_URL || "http://localhost:8080/api/v1",
});

export const setAuthToken = (token?: string) => {
    if (token) {
        api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    } else {
        delete api.defaults.headers.common["Authorization"];
    }
}

export default axios;
import React from 'react'
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import {
    useForm,
    Controller,
    SubmitHandler,
    useFormState,
  } from "react-hook-form";
  
  import {
    firstNameValidation,
    lastNameValidation,
    passportNumberValidation,
    phoneValidation,
  } from "./validate";
  
  import PassengersList from "./PassengersList";
  import FlightDetails from "./FlightDetails";

type Props = {}

interface IBookingForm {
    firstName: string;
    lastName: string;
    birthdate: Date;
    passportNumber: number;
    passportCountry: string;
    phone: number;
  }

const BookingConf = (props: Props) => {
    const { handleSubmit, control } = useForm<IBookingForm>();
  const { errors } = useFormState({ control });
  const onSubmit: SubmitHandler<IBookingForm> = (data) => console.log(data);
  return (
    <div><FlightDetails />

    <form onSubmit={handleSubmit(onSubmit)}>
      <Typography sx={{ mt: 5 }} variant="subtitle1">
        Passenger details
      </Typography>
      <Grid
        container
        rowSpacing={1}
        columnSpacing={{ xs: 1, sm: 2, md: 3 }}
        justifyContent="center"
        alignItems="center"
      >
        <Grid item xs={4}>
          <Controller
            control={control}
            rules={firstNameValidation}
            name="firstName"
            render={({ field }) => (
              <TextField
                size="small"
                label="FirstName"
                margin="normal"
                fullWidth={true}
                helperText={errors.firstName?.message}
                error={!!errors.firstName?.message}
                onChange={(e) => field.onChange(e)}
                value={field.value}
              />
            )}
          />
        </Grid>
        <Grid item xs={4}>
          <Controller
            control={control}
            rules={lastNameValidation}
            name="lastName"
            render={({ field }) => (
              <TextField
                label="LastName"
                size="small"
                margin="normal"
                fullWidth={true}
                helperText={errors.lastName?.message}
                error={!!errors.lastName?.message}
                onChange={(e) => field.onChange(e)}
                value={field.value}
              />
            )}
          />
        </Grid>
        <Grid item xs={4}>
          <Controller
            control={control}
            name="birthdate"
            render={({ field }) => (
              <TextField
                size="small"
                margin="normal"
                fullWidth={true}
                label="Birthdate"
                type="date"
                defaultValue="2017-05-24"
                value={field.value}
              />
            )}
          />
        </Grid>
        <Grid item xs={4}>
          <Controller
            control={control}
            rules={passportNumberValidation}
            name="passportNumber"
            render={({ field }) => (
              <TextField
                label="Passport Number"
                size="small"
                margin="normal"
                fullWidth={true}
                helperText={errors.passportNumber?.message}
                error={!!errors.passportNumber?.message}
                onChange={(e) => field.onChange(e)}
                value={field.value}
              />
            )}
          />
        </Grid>
        <Grid item xs={4}>
          <Controller
            control={control}
            // rules={loginValidation}
            name="passportCountry"
            render={({ field }) => (
              <FormControl fullWidth size="small" margin="normal">
                <InputLabel id="demo-simple-select-label">
                  Passport country
                </InputLabel>
                <Select
                  id="demo-simple-select"
                  value={field.value}
                  label="Passport Country"
                  onChange={(e) => field.onChange(e)}
                >
                  <MenuItem value={"Russia"}>Russia</MenuItem>
                  <MenuItem>...</MenuItem>
                  <MenuItem>...</MenuItem>
                </Select>
              </FormControl>
            )}
          />
        </Grid>
        <Grid item xs={4}>
          <Controller
            control={control}
            rules={phoneValidation}
            name="phone"
            render={({ field }) => (
              <TextField
                label="Phone"
                size="small"
                margin="normal"
                fullWidth={true}
                helperText={errors.phone?.message}
                error={!!errors.phone?.message}
                onChange={(e) => field.onChange(e)}
                value={field.value}
              />
            )}
          />
        </Grid>
      </Grid>
      <Button
        type="submit"
        variant="contained"
        className="btn-add"
        fullWidth={true}
        disableElevation={true}
        sx={{ width: "25%", mt: 2 }}
      >
        Add passenger
      </Button>
    </form>

    <PassengersList /></div>
  )
}

export default BookingConf
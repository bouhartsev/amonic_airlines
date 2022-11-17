import Typography from '@mui/material/Typography';
import TextField, { TextFieldProps } from '@mui/material/TextField';
import Button from '@mui/material/Button';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';

import React, { FC, useState } from 'react';
import {
  useForm,
  Controller,
  SubmitHandler,
  useFormState,
} from 'react-hook-form';

import {
  firstNameValidation,
  lastNameValidation,
  passportNumberValidation,
  phoneValidation,
} from './validate';

import './index.css';
import { PassengersList } from './PassengersList';
import { FlightDetails } from './FlightDetails';

interface IBookingForm {
  firstName: string;
  lastName: string;
  birthdate: Date;
  passportNumber: number;
  passportCountry: string;
  phone: number;
}

export const BookingPage: FC = () => {
  const { handleSubmit, control } = useForm<IBookingForm>();
  const { errors } = useFormState({ control });
  const onSubmit: SubmitHandler<IBookingForm> = (data) => console.log(data);

  return (
    <div className="booking-form">
      <Typography variant="h4" sx={{ textAlign: 'center', marginBottom: 5 }}>
        Booking confirmation
      </Typography>

      <FlightDetails />

      <form className="booking-form__form" onSubmit={handleSubmit(onSubmit)}>
        <p className="booking-form__title">Passenger details</p>
        <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 2, md: 3 }}>
          <Grid item xs={4}>
            <Controller
              control={control}
              rules={firstNameValidation}
              name="firstName"
              render={({ field }) => (
                <TextField
                  label="FirstName"
                  size="small"
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
                  label="Birthdate"
                  type="date"
                  defaultValue="2017-05-24"
                  sx={{ width: 220 }}
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
                <FormControl fullWidth>
                  <InputLabel id="demo-simple-select-label">
                    Passport country
                  </InputLabel>
                  <Select
                    id="demo-simple-select"
                    value={field.value}
                    label="Passport Country"
                    onChange={(e) => field.onChange(e)}
                  >
                    <MenuItem value={'Russia'}>Russia</MenuItem>
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
          sx={{
            marginTop: 2,
          }}
        >
          Add passenger
        </Button>
      </form>

      <PassengersList />
    </div>
  );
};

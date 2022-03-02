<template>
  <v-container>
    <v-row justify="center">
      <div class="text-center">
        <b>Set/Delete Available Time for Your Patients</b>
      </div>
      <br />
      <br />
    </v-row>
    <v-row>
      <v-calendar ref="calendar"></v-calendar>
    </v-row>
    <v-row align="center">
      <v-time-picker v-model="picker" ampm-in-title></v-time-picker>
      <v-spacer />
      <v-col>
        <v-btn large class="mr-2" color="primary"
          >Add new available time slot <v-icon large>mdi-bell-plus</v-icon></v-btn
        >
        <v-btn large color="red"
          >Cancel chosen available time slots
          <v-icon large>mdi-delete-clock</v-icon></v-btn
        >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'SetAvailablePage',
  layout: 'dash',
  middleware({ store, redirect }) {
    if (store.state.auth.user.profile.role !== 'Doctor') {
      return redirect('/dashboard')
    }
  },
}
</script>

<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline"> Register </v-card-title>
      </v-card>
      <v-card>
          <RegisterForm />
          <br />
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'RegisterPage',
  middleware({ store, redirect }) {
    if (store.state.auth.user) {
      return redirect('/dashboard')
    }
  },
}
</script>

<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline"> Sign In </v-card-title>
      </v-card>
      <v-card>
          <LoginForm />
          <br />
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'LoginPage',
  middleware({ store, redirect }) {
    if (store.state.auth.user) {
      return redirect('/dashboard')
    }
  },
}
</script>

<template>
  <form class="ma-3">
    <v-snackbar v-model="snackbar" color="red" centered transition="slide-x-transition">
      {{ snackbarText }}
      <template #action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <v-text-field
      v-model.trim="email"
      :error-messages="emailErrors"
      label="E-mail"
      required
      @input="$v.email.$touch()"
      @blur="$v.email.$touch()"
    ></v-text-field>
    <v-text-field
      v-model.trim="password"
      type="password"
      :error-messages="passwordErrors"
      label="Password"
      required
      @input="$v.password.$touch()"
      @blur="$v.password.$touch()"
    ></v-text-field>
    <v-btn class="mr-4" color="primary" @click="submit"> submit </v-btn>
    <v-btn @click="clear"> clear </v-btn>
    <span class="mr-3" />
    <v-progress-circular
      v-show="loading"
      indeterminate
      color="red"
    ></v-progress-circular>
  </form>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, email } from 'vuelidate/lib/validators'

export default {
  mixins: [validationMixin],

  validations: {
    email: { required, email },
    password: { required },
  },

  data: () => ({
    snackbar: false,
    snackbarText: '',
    email: '',
    password: '',
    loading: false,
  }),

  computed: {
    emailErrors() {
      const errors = []
      if (!this.$v.email.$dirty) return errors
      !this.$v.email.email && errors.push('Must be valid e-mail')
      !this.$v.email.required && errors.push('E-mail is required')
      return errors
    },
    passwordErrors() {
      const errors = []
      if (!this.$v.password.$dirty) return errors
      !this.$v.password.required && errors.push('Password is required')
      return errors
    },
  },

  methods: {
    async submit() {
      this.$v.$touch()
      if (! this.$v.$invalid) {
        try {
          this.loading = true
          const res = await this.$axios.$post('auth/login', {
            email: this.email,
            password: this.password
          })
          this.$store.commit('auth/login', res.data);
          await this.$router.push({ path: '/dashboard' })
        } catch (err) {
          this.snackbarText = "Your password or e-mail is incorrect. Please try again..."
          this.snackbar = true
        }

        this.loading = false
      }
    },
    clear() {
      this.$v.$reset()
      this.email = ''
      this.password = ''
    },
  },
}
</script>

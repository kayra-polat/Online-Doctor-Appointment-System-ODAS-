<template>
  <v-app dark>
    <v-app-bar fixed app>
      <v-icon class="ml-4" x-large>mdi-shield-crown-outline</v-icon>
      <v-toolbar-title class="ml-2" v-text="title" />
      <v-spacer />
      <v-icon large>mdi-login</v-icon>LOGIN
      <span class="ml-4"></span>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <v-spacer />
      <span>ODAS LLC. &copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'AdminLayout',
  data () {
    return {
      title: 'ADMIN INTERFACE'
    }
  }  
}
</script>

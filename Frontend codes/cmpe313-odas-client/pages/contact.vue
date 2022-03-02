<template>
  <v-row>
    <v-col class="text-center">
      <br />
      <span><b>CONTACT FORM</b></span>
      <form class="ma-3">
        <v-text-field
          label="E-mail"
        ></v-text-field>
        <v-text-field
          label="Name"
        ></v-text-field>
        <v-textarea
          label="Type your message here..." 
        ></v-textarea>
        <v-btn color="primary"> send </v-btn>
      </form>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'ContactPage',
  layout: 'dash',
}
</script>

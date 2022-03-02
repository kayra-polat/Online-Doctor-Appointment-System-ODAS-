<template>
  <v-col>
    <v-row justify="center" align-content="center" align-text="center">
      <v-col>
        <v-btn class="mt-16" x-large>
          Manage Database
          <v-icon x-large>mdi-database-cog</v-icon>
        </v-btn>
      </v-col>
      <v-col>
        <v-btn class="mt-16" x-large>
          Show logs
          <v-icon x-large>mdi-file-find</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center" align-content="center" align-text="center">
      <v-col>
        <v-btn class="mt-16" x-large>
          Change System/Application Settings
          <v-icon x-large>mdi-cog</v-icon>
        </v-btn>
      </v-col>
      <v-col>
        <v-btn class="mt-16" x-large>
          Statistics
          <v-icon x-large>mdi-chart-areaspline</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center" align-content="center" align-text="center">
      <v-col>
        <v-btn class="mt-16" x-large color="red">
          <v-icon x-large>mdi-console</v-icon>
          Open CLI
        </v-btn>
      </v-col>
    </v-row>
  </v-col>
</template>

<script>
export default {
  name: 'AdminPage',
  layout: 'admin',
}
</script>

<template>
  <b-container>
    <p>Hello, Bootstrap-Vue</p>
    <b-btn 
      variant="primary"
      @click="showAlert = true"
    >
      Click
    </b-btn>
    <b-alert 
      v-bind:show="showAlert"
      variant="success"
    >
      You clicked the button!
    </b-alert>
  </b-container>
</template>


<script>
export default {
  data: () => ({
    showAlert: false
  })
};
</script>
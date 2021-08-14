<template>
  <div class="hello">
    <h1>This is my IPFS API</h1>

    <!-- CREATE AN INPUT BOX TO ADD IN YOUR HASH -->
    <!-- Get data through `POST`, encrypt it, and then store it in IPFS -->
    <!-- Through a `GET`, retrieve CID from IPFS, decrypt it, and then serve the content as a reply -->

    <!-- Using https://bootstrap-vue.org/ build a simple UI that interacts with the API you’ve built to store and retrieve encrypted data from IPFS: -->
    <!-- The UI must ask for the encryption key as soon as it’s loaded -->
    <!-- Use a text field for data -->
    <!-- As a challenge you can also implement a file upload feature -->
    <br />
    <br />

    <b-form @submit="onSubmit1" @reset="onReset1" v-if="show">
      <b-form-group
        id="input-group-1"
        label="Send a String and Get a Hash!"
        label-for="input-1"
        description="We'll never share your hash with anyone else."
      >
        <!-- STRING SENDER -->
        <b-form-textarea
          id="input-1"
          v-model="form1.data"
          placeholder="Enter your String..."
          rows="3"
          max-rows="6"
          required
        ></b-form-textarea>
      </b-form-group>

      <!-- SUBMIT BUTTONS -->
      <b-container style="display: flex; flex-direction: row; width: 30%; justify-content: space-between; align-items: center;">
        <b-button type="submit" variant="primary">Submit</b-button>
        <b-button type="reset" variant="danger">Reset</b-button>
      </b-container>

    </b-form>
    <br>
    <br>

    <!-- <b-card class="mt-3" header="Form Data Result">
      <pre class="m-0">{{ form1 }}</pre>
    </b-card> -->

    <h3>{{ stringHash }}</h3>

    <br>
    <br>
    <b-form @submit="onSubmit2" @reset="onReset2" v-if="show">
      <b-form-group
        id="input-group-2"
        label="Get Hash Address:"
        label-for="input-2"
        description="We'll never share your hash with anyone else."
      >
        <!-- HASH GRABBER -->
        <b-form-textarea
          id="input-2"
          v-model="form2.hash"
          placeholder="Enter your Hash..."
          rows="3"
          max-rows="6"
          required
        ></b-form-textarea>
      </b-form-group>

      <!-- SUBMIT BUTTONS -->
      <b-container style="display: flex; flex-direction: row; width: 30%; justify-content: space-between; align-items: center;">
        <b-button type="submit" variant="primary">Submit</b-button>
        <b-button type="reset" variant="danger">Reset</b-button>
      </b-container>

    </b-form>

    <!-- <b-card class="mt-3" header="Form Data Result">
      <pre class="m-0">{{ form2 }}</pre>
    </b-card> -->

    <pre class="mt-3 mb-0">{{ text }}</pre>

    
    <!-- {{}} output object attributes and function return values   -->
    <!-- <h3>{{ details() }}</h3> -->
    <br />
    <h3>{{ albums }}</h3>
    <!-- <h1 v-for="data in ydata" v-bind:key="data">{{ data }}</h1>
    <h3 v-for="item in xdata" v-bind:key="item">{{ item }}</h3> -->
  </div>
</template>
<script>
import axios from "axios";
export default {
  name: "apidata",
  //data used to define the properties of the returned data
  data() {
    return {
      showAlert: false,
      albums: null,
      stringHash: null,
      form1: {
        data: "",
      },
      form2: {
        hash: "",
      },
      show: true,
    };
  },
  //   @ A method for defining the js
  methods: {
    onSubmit1(event) {
      event.preventDefault();
      alert(JSON.stringify(this.form1.data));

      axios
        .post(`http://localhost:8000/${this.form1.data}`)
        // .get("http://localhost:8000/"+JSON.stringify(this.form.hash))
        .then(
          (response) => (this.stringHash = response.data)
          // (this.xdata = response.data.legend_data),
          // (this.ydata = response.data.xAxis_data)
        )
        .catch((error) => console.log("Turn on the Gin API!", error));
    },
    onReset1(event) {
      event.preventDefault();
      // Reset our form values
      this.form1.data = "";
      // Trick to reset/clear native browser form validation state
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    },
    onSubmit2(event) {
      event.preventDefault();
      alert(JSON.stringify(this.form2.hash));

      axios
        .get(`http://localhost:8000/${this.form2.hash}`)
        // .get("http://localhost:8000/"+JSON.stringify(this.form.hash))
        .then(
          (response) => (this.albums = response.data)
          // (this.xdata = response.data.legend_data),
          // (this.ydata = response.data.xAxis_data)
        )
        .catch((error) => console.log("Turn on the Gin API!", error));
    },
    onReset2(event) {
      event.preventDefault();
      // Reset our form values
      this.form2.hash = "";
      // Trick to reset/clear native browser form validation state
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    },
    details: function () {
      return this.site;
    },
  },
  // mounted() {
  //   //response returns a json { "data": "Data", "status": "status code", "statusText": "Status text", "headers": { "content-type": "application/json; charset = utf-8 "}," config ":" profile "," method ":" method "," url ":" request url "," request ":" request body "}
  //   axios
  //     .get("http://localhost:8000/albums")
  //     // .get("http://localhost:8000/"+JSON.stringify(this.form.hash))
  //     .then(
  //       (response) => (this.albums = response.data)
  //       // (this.xdata = response.data.legend_data),
  //       // (this.ydata = response.data.xAxis_data)
  //     )
  //     .catch((error) => console.log("Turn on the Gin API!", error));
  // },
};
</script>
<style>
/* [Use of css class selector [multiple styles into force priority] */
.hello {
  font-weight: normal;
  text-align: center;
  font-size: 8pt;
  padding: 8rem;
}
h3 {
  text-align: center;
  font-size: 20pt;
  color: red;
}
</style>
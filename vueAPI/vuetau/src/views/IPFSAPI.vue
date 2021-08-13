<template>
  <div class="hello">
    <!-- {{}} output object attributes and function return values   -->
    <h1>{{ msg }}</h1>
    <h1>site: {{ site }}</h1>
    <h1>url: {{ url }}</h1>
    <h3>{{ details() }}</h3>
    <h3>{{ albums }}</h3>
    <h1 v-for="data in ydata" v-bind:key="data">{{ data }}</h1>
    <h3 v-for="item in xdata" v-bind:key="item">{{ item }}</h3>
  </div>
</template>
<script>
import axios from "axios";
export default {
  name: "apidata",
  //data used to define the properties of the returned data
  data() {
    return {
      msg: "hello, xxbandy! ",
      site: "bgops",
      url: "https://xxbandy.github.io",
      xdata: null,
      ydata: null,
      albums: null
    };
  },
  //   @ A method for defining the js
  methods: {
    details: function () {
      return this.site;
    },
  },
  mounted() {
    //response returns a json { "data": "Data", "status": "status code", "statusText": "Status text", "headers": { "content-type": "application/json; charset = utf-8 "}," config ":" profile "," method ":" method "," url ":" request url "," request ":" request body "}
    axios
      .get("http://localhost:8000/albums")
      .then(
        (response) => (
          (this.albums = response.data)
          // (this.xdata = response.data.legend_data),
          // (this.ydata = response.data.xAxis_data)
        )
      );
  },
};
</script>
<style>
/* [Use of css class selector [multiple styles into force priority] */
.hello {
  font-weight: normal;
  text-align: center;
  font-size: 8pt;
}
h3 {
  text-align: center;
  font-size: 20pt;
  color: red;
}
</style>
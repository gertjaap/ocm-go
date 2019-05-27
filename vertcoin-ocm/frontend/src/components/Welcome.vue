<template>
  <div class="container">
    <p>Welcome to the Vertcoin One Click Miner</p>
    <div class="col-6">
    	<p>Enter a password to protect your proceeds. <b>Don't lose it</b>.</p>
	    <p><input type="password" placeholder="Password" /><br/>
	    <input type="password" placeholder="Confirm Password" /></p>
		<p>
	    <a @click="start">Start Mining!</a>
		</p>
    </div>
    <div class="col-6">
    	<p><b>Your Graphics Card(s):</b></p>
		<p>NVidia GTX 1080ti</p>
		<p>Expected proceeds: <b>0.98 VTC</b> per day</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: " "
    };
  },
  methods: {
    start: function() {
      /*var self = this;
      window.backend.basic().then(result => {
        self.message = result;
      });*/
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 1em;
  position: relative;
  min-height: 2rem;
  width: 100%;
}
input {
  font-size: 1em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 1px solid white;
  border-radius: 5px;
  padding: 4px;
  cursor: pointer;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>

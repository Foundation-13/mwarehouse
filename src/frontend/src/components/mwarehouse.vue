<style>
  input[type="file"]{
    position: absolute;
    top: -500px;
  }
  div.file-listing img{
    max-width: 90%;
  }
</style>

<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label>Files
        <input type="file" id="files" ref="files" accept="image/*" multiple v-on:change="handleFilesUpload()"/>
      </label>
    </div>
    <div class="large-12 medium-12 small-12 cell">
      <div class="grid-x">
        <div v-for="(file, key) in files" :key="file" class="large-4 medium-4 small-6 cell file-listing">
          {{ file.name }}
          <img class="preview" v-bind:ref="'image'+parseInt( key )"/>
        </div>
      </div>
    </div>
    <br>
    <div class="large-12 medium-12 small-12 cell">
      <button v-on:click="addFiles()">Add Files</button>
    </div>
    <br>
    <div class="large-12 medium-12 small-12 cell">
      <button v-on:click="submitFiles()">Submit</button>
    </div>
  </div>
</template>

<script>
  import axios from "axios"
  export default {
    /*
      Defines the data used by the component
    */
    data(){
      return {
        files: [],
        myFile: ''
      }
    },

    /*
      Defines the method used by the component
    */
    methods: {
      /*
        Adds a file
      */
      addFiles(){
        this.$refs.files.click();
      },

      /*
        Submits files to the server
      */
      submitFiles(){
        /*
          Initialize the form data
        */
        let formData = new FormData();

        /*
          Iteate over any file sent over appending the files
          to the form data.
        */
        formData.append('file', this.myFile)
        /*
          Make the request to the POST /select-files URL
        */
        axios.post('http://localhost:8765/media',
          formData,
          {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
          }
        ).then(function(){
          console.log('SUCCESS!!');
        })
        .catch(function(){
          console.log('FAILURE!!');
        });
      },

      /*
        Handles the uploading of files
      */
      handleFilesUpload(){
        let uploadedFiles = this.$refs.files.files;
        this.myFile = uploadedFiles[0];
        /*
          Adds the uploaded file to the files array
        */
        for( var i = 0; i < uploadedFiles.length; i++ ){
          this.files.push( uploadedFiles[i] );
        }
        this.getImagePreviews();
      },

      /*
        Removes a select file the user has uploaded
      */
      getImagePreviews(){
        for( let i = 0; i < this.files.length; i++ ){
          if ( /\.(jpe?g|png|gif)$/i.test( this.files[i].name ) ) {
            let reader = new FileReader();


            reader.addEventListener("load", function(){
              this.$refs['image'+parseInt( i )][0].src = reader.result;
            }.bind(this), false);

           
            reader.readAsDataURL( this.files[i] );
          }
        }
      }
    }
  }
</script>
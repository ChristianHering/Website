<template>
    <div class="col">
        <div v-html="html" />
        <input v-model="html">
        <button v-on:click="postCreate(html)">Post the post!</button>
        <button v-on:click="stagedPostCreate(html)">Save for later</button>
    </div>
</template>

<script>
    import axios from "axios"; //For REST Requests

    export default {
        name: "PostCreate",
        data: function() {
            return {
                html: "",
                success: null,
                failure: null
            }
        },
        methods: {
            postCreate: function(error) {
                axios.post(
                    "http://admin.christianhering.com/postCreate",
                    error
                )
                .then((response) => {
                    console.log(response);
                }, (error) => {
                    console.log(error);
                });
            },
            stagedPostCreate: function(error) {
                axios.post(
                    "http://admin.christianhering.com/postCreate",
                    error
                )
                .then((response) => {
                    console.log(response);
                }, (error) => {
                    console.log(error);
                });
            }
        }
    };
</script>
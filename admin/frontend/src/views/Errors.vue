<template>
    <div class="errors">
        <div v-if="error.length > 0">
            <h1>Error ID: {{error[0].ID}}</h1>
            <h3>{{error[0].Date}}</h3>
            <h3>{{error[0].Host}}{{error[0].URL}}</h3>
            <h5>{{error[0].Error}}</h5>
            <div class="row">
                <button v-on:click="deleteError(error[0])">Delete Error</button>
                <button v-on:click="deleteErrorType(error[0])"> Delete Type</button>
            </div>
        </div>
        <div v-else> <!-- Display a generic message if there's no errors -->
            <h1>All Clear!</h1>
        </div>

        <div v-if="error.length > 1">
            <h1>Error ID: {{error[1].ID}}</h1>
            <h3>{{error[1].Date}}</h3>
            <h3>{{error[1].Host}}{{error[1].URL}}</h3>
            <h5>{{error[1].Error}}</h5>
            <div class="row">
                <button v-on:click="deleteError(error[1])">Delete Error</button>
                <button v-on:click="deleteErrorType(error[1])"> Delete Type</button>
            </div>
        </div>

        <div v-if="error.length > 2">
            <h1>Error ID: {{error[2].ID}}</h1>
            <h3>{{error[2].Date}}</h3>
            <h3>{{error[2].Host}}{{error[2].URL}}</h3>
            <h5>{{error[2].Error}}</h5>
            <div class="row">
                <button v-on:click="deleteError(error[2])">Delete Error</button>
                <button v-on:click="deleteErrorType(error[2])"> Delete Type</button>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from "axios"; //For REST Requests

    export default {
        name: "Errors",
        data: function() {
            return {error: []}
        },
        mounted: function() { //"Main"
            this.refreshErrorList()
        },
        methods: {
            refreshErrorList: function() {
                axios.get("http://admin.christianhering.com/error")
                    .then((response) => {
                        this.error = response.data;
                    })
                    .catch(
                        error => console.log(error)
                    );
            },
            deleteError: function(error) {
                axios.post(
                    "http://admin.christianhering.com/errorDelete",
                    error
                )
                .then((response) => {
                    console.log(response);
                }, (error) => {
                    console.log(error);
                });

                this.refreshErrorList();
            },
            deleteErrorType: function(error) {
                axios.post(
                    "http://admin.christianhering.com/errorDeleteType",
                    error
                )
                .then((response) => {
                    console.log(response);
                }, (error) => {
                    console.log(error);
                });

                this.refreshErrorList();
            }
        }
    };
</script>
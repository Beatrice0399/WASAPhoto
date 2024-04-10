<script>
export default {
    data(){
        return{
            photoPATH: "",
            liked: false,
            comments: [],
            likes: [],
        }
    },
    props: ["phid", "uid", "comments", "likes", "date"],

    methods: {
        loadPhoto() {

        }
    }
}
</script>
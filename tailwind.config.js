/** @type {import('tailwindcss').Config} */
export default {
    content: ["./view/**/*.tmpl"],
    theme: {
        extend: {
            flex: {
                '2': '2 2 0%',
                '3': '3 3 0%',
                '4': '4 4 0%',
            },
        },
    },
    plugins: [],
}



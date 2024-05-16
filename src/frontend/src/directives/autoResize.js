// src/directives/autoResize.js
export default {
    mounted(el) {
        el.style.height = "auto";
        el.style.height = el.scrollHeight + "px";
        el.addEventListener("input", () => {
            el.style.height = "auto";
            el.style.height = el.scrollHeight + "px";
        });
    },
    updated(el) {
        el.style.height = "auto";
        el.style.height = el.scrollHeight + "px";
    }
};

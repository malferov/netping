export const isDev = process.env.NODE_ENV !== "production";

class DevLogger {
    constructor() {
        this.colors = {notify: '#006400', warning: '#808000', error: 'red'};
        this.style = `
            padding: 3px 15px;
            font-size: 14px;
            color: #fff;
            font-weight: bold;
            border-radius: 10px;
        `;
    }

    logInProduction() {
        if(!isDev) console.log(...arguments);
    }

    notify() {
        this.log("%cNotify", `${this.style}background: ${this.colors.notify};` , ...arguments);
    }

    warning() {
        this.log("%cWarning", `${this.style}background-color: ${this.colors.warning};`, ...arguments);
    }

    error() {
        this.log("%cError", `${this.style}background-color: ${this.colors.error};`, ...arguments);
    }

    handler(handler) {
        if (isDev) handler();
    }

    log() {
        if (isDev) console.log(...arguments);
    }
}
export default new DevLogger();

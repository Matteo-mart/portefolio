const i18n = createI18n({
  legacy: false, 
  locale: 'fr',  
  messages: {
    fr: {
        // ex
    //   title: "Mes TodoLists",
    // date:
    // description:
    // technologie:
    // explication:
    // probleme:
    // solution:
    // url_source:


    },
    en: {
        // ex
    //   todolist: "My TodoLists",
    }
  }


});

createApp(App).use(i18n).mount('#app');

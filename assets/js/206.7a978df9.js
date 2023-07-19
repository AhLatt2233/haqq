(window.webpackJsonp=window.webpackJsonp||[]).push([[206],{760:function(e,t,a){"use strict";a.r(t);var s=a(1),o=Object(s.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"run-a-node"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#run-a-node"}},[e._v("#")]),e._v(" Run a Node")]),e._v(" "),a("p",{attrs:{synopsis:""}},[e._v("Configure and run an Haqq node")]),e._v(" "),a("h2",{attrs:{id:"pre-requisite-readings"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#pre-requisite-readings"}},[e._v("#")]),e._v(" Pre-requisite Readings")]),e._v(" "),a("ul",[a("li",{attrs:{prereq:""}},[a("RouterLink",{attrs:{to:"/quickstart/installation.html"}},[e._v("Installation")])],1),e._v(" "),a("li",{attrs:{prereq:""}},[a("RouterLink",{attrs:{to:"/quickstart/binary.html"}},[a("code",[e._v("haqqd")])])],1)]),e._v(" "),a("h2",{attrs:{id:"automated-deployment"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#automated-deployment"}},[e._v("#")]),e._v(" Automated deployment")]),e._v(" "),a("p",[e._v("Run the local node by running the "),a("code",[e._v("init.sh")]),e._v(" script in the base directory of the repository.")]),e._v(" "),a("div",{staticClass:"custom-block warning"},[a("p",[e._v("The script below will remove any pre-existing binaries installed. Use the manual deploy if you want\nto keep your binaries and configuration files.")])]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Li9pbml0LnNoCg=="}}),e._v(" "),a("h2",{attrs:{id:"manual-deployment"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#manual-deployment"}},[e._v("#")]),e._v(" Manual deployment")]),e._v(" "),a("p",[e._v("The instructions for setting up a brand new full node from scratch are the the same as running a\n"),a("RouterLink",{attrs:{to:"/guides/localnet/single_node.html#manual-localnet"}},[e._v("single node local testnet")]),e._v(".")],1),e._v(" "),a("h2",{attrs:{id:"start-node"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#start-node"}},[e._v("#")]),e._v(" Start node")]),e._v(" "),a("p",[e._v("To start your node, just type:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"aGFxcWQgc3RhcnQgLS1qc29uLXJwYy5lbmFibGU9dHJ1ZSAtLWpzb24tcnBjLmFwaT0mcXVvdDtldGgsd2ViMyxuZXQmcXVvdDsK"}}),e._v(" "),a("p",[e._v("Alternatively, you can run your node as a service. To do this, first create systemd service file in /etc/systemd/system as in the following example:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"dGVlICRIT01FL2hhcXFkLnNlcnZpY2UgJmd0OyAvZGV2L251bGwgJmx0OyZsdDtFT0YKW1VuaXRdCkRlc2NyaXB0aW9uPUhhcXEgTm9kZQpBZnRlcj1uZXR3b3JrLnRhcmdldApbU2VydmljZV0KVXNlcj1yb290ClR5cGU9c2ltcGxlCkV4ZWNTdGFydD0vcm9vdC9nby9iaW4vaGFxcWQgc3RhcnQKUmVzdGFydD1vbi1mYWlsdXJlClJlc3RhcnRTZWM9MTAKTGltaXROT0ZJTEU9NjU1MzUKW0luc3RhbGxdCldhbnRlZEJ5PW11bHRpLXVzZXIudGFyZ2V0CkVPRgoKc3VkbyBtdiAkSE9NRS9oYXFxZC5zZXJ2aWNlIC9ldGMvc3lzdGVtZC9zeXN0ZW0vCg=="}}),e._v(" "),a("p",[e._v("Next, enable the haqqd service, reload systemd, and start the node:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c3VkbyBzeXN0ZW1jdGwgZW5hYmxlIGhhcXFkCnN1ZG8gc3lzdGVtY3RsIGRhZW1vbi1yZWxvYWQKc3VkbyBzeXN0ZW1jdGwgc3RhcnQgaGFxcWQK"}}),e._v(" "),a("p",[e._v("Check that your node is working correctly in logs:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"am91cm5hbGN0bCAtdSBoYXFxZCAtZiAtbyBjYXQK"}}),e._v(" "),a("h2",{attrs:{id:"key-management"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#key-management"}},[e._v("#")]),e._v(" Key Management")]),e._v(" "),a("p",[e._v("To run a node with the same key every time: replace "),a("code",[e._v("haqqd keys add $KEY")]),e._v(" in "),a("code",[e._v("./init.sh")]),e._v(" with:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"ZWNobyAmcXVvdDt5b3VyIG1uZW1vbmljIGhlcmUmcXVvdDsgfCBoYXFxZCBrZXlzIGFkZCAkS0VZIC0tcmVjb3Zlcgo="}}),e._v(" "),a("div",{staticClass:"custom-block tip"},[a("p",[e._v("Haqq currently only supports 24 word mnemonics.")])]),e._v(" "),a("p",[e._v("You can generate a new key/mnemonic with:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"aGFxcWQga2V5cyBhZGQgJEtFWQo="}}),e._v(" "),a("p",[e._v("To export your haqq key as an Ethereum private key (for use with "),a("RouterLink",{attrs:{to:"/guides/keys-wallets/metamask.html"}},[e._v("Metamask")]),e._v(" for example):")],1),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"aGFxcWQga2V5cyB1bnNhZmUtZXhwb3J0LWV0aC1rZXkgJEtFWQo="}}),e._v(" "),a("p",[e._v("For more about the available key commands, use the "),a("code",[e._v("--help")]),e._v(" flag")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"aGFxcWQga2V5cyAtaAo="}}),e._v(" "),a("h3",{attrs:{id:"keyring-backend-options"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#keyring-backend-options"}},[e._v("#")]),e._v(" Keyring backend options")]),e._v(" "),a("p",[e._v("The instructions above include commands to use "),a("code",[e._v("test")]),e._v(" as the "),a("code",[e._v("keyring-backend")]),e._v(". This is an unsecured\nkeyring that doesn't require entering a password and should not be used in production. Otherwise,\nHaqq supports using a file or OS keyring backend for key storage. To create and use a file\nstored key instead of defaulting to the OS keyring, add the flag "),a("code",[e._v("--keyring-backend file")]),e._v(" to any\nrelevant command and the password prompt will occur through the command line. This can also be saved\nas a CLI config option with:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"aGFxcWQgY29uZmlnIGtleXJpbmctYmFja2VuZCBmaWxlCg=="}}),e._v(" "),a("div",{staticClass:"custom-block tip"},[a("p",[e._v("For more information about the Keyring and its backend options, click "),a("RouterLink",{attrs:{to:"/guides/keys-wallets/keyring.html"}},[e._v("here")]),e._v(".")],1)]),e._v(" "),a("h2",{attrs:{id:"clearing-data-from-chain"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#clearing-data-from-chain"}},[e._v("#")]),e._v(" Clearing data from chain")]),e._v(" "),a("h3",{attrs:{id:"reset-data"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#reset-data"}},[e._v("#")]),e._v(" Reset Data")]),e._v(" "),a("p",[e._v("Alternatively, you can "),a("strong",[e._v("reset")]),e._v(" the blockchain database, remove the node's address book files, and reset the "),a("code",[e._v("$HOME/.haqqd/data/priv_validator_state.json")]),e._v(" to the genesis state.")]),e._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",[e._v("If you are running a "),a("strong",[e._v("validator node")]),e._v(", always be careful when doing "),a("code",[e._v("haqqd unsafe-reset-all")]),e._v(". You should never use this command if you are not switching "),a("code",[e._v("chain-id")]),e._v(".")])]),e._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",[a("strong",[e._v("IMPORTANT")]),e._v(": Make sure that every node has a unique file "),a("code",[e._v("priv_validator_key.json")]),e._v(". "),a("strong",[e._v("Do not")]),e._v(" copy the "),a("code",[e._v("priv_validator_key.json")]),e._v(" from an old node to multiple new nodes. Running two nodes with the same "),a("code",[e._v("priv_validator_key.json")]),e._v(" will cause you to double sign!")])]),e._v(" "),a("p",[e._v("First, remove the outdated files and reset the data.")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c3VkbyBzeXN0ZW1jdGwgc3RvcCBoYXFxZApybSAkSE9NRS8uaGFxcWQvY29uZmlnL2FkZHJib29rLmpzb24gJEhPTUUvLmhhcXFkL2NvbmZpZy9nZW5lc2lzLmpzb24KaGFxcWQgdW5zYWZlLXJlc2V0LWFsbAo="}}),e._v(" "),a("p",[e._v("Your node is now in a pristine state while keeping the original "),a("code",[e._v("priv_validator_key.json")]),e._v(" and "),a("code",[e._v("config.toml")]),e._v(". If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.")]),e._v(" "),a("h3",{attrs:{id:"delete-data"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#delete-data"}},[e._v("#")]),e._v(" Delete Data")]),e._v(" "),a("p",[e._v("Data for the "+e._s(e.$themeConfig.project.binary)+" binary should be stored at "),a("code",[e._v("~/."+e._s(e.$themeConfig.project.binary))]),e._v(", respectively by default. To "),a("strong",[e._v("delete")]),e._v(" the existing binaries and configuration, run:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"cm0gLXJmIH4vLmhhcXFkCg=="}}),e._v(" "),a("p",[e._v("To clear all data except key storage (if keyring backend chosen) and then you can rerun the full node installation commands from above to start the node again.")]),e._v(" "),a("h2",{attrs:{hide:"",id:"next"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#next"}},[e._v("#")]),e._v(" Next")]),e._v(" "),a("p",{attrs:{hide:""}},[e._v("Learn about running a Haqq "),a("RouterLink",{attrs:{to:"/testnet/"}},[e._v("testnet")])],1)],1)}),[],!1,null,null,null);t.default=o.exports}}]);
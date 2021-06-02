const enableEthereum = async () => {
  if (window.ethereum) {
    // await ethereum.request({ method: "eth_accounts" });
    await ethereum.enable();
  }

  // Check wether it's already injected by something else (like Metamask or Parity Chrome plugin)
  if (typeof web3 !== "undefined") {
    console.log("Metamask detected");
    web3 = new Web3(web3.currentProvider);  
  } else {
    // Or connect to a node
    web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8080"));
  }

  // set publicAdress
  var publicAddress = await getMetamaskAddress();
  console.log(publicAddress);
  document.getElementById("address").value = publicAddress;
};

const getMetamaskAddress = async () => {
  // Log all eth accounts
  const accounts = await ethereum.request({ method: "eth_accounts" });
  console.log("Accounts are", accounts);

  // Grab the metamask public address
  return accounts[0];
};

const getNonce = async () => {
  console.log("GET nonce called");
  try {
    const response = await axios.post("/request_nonce");
    console.log(response);

    document.getElementById("nonce").value = response.data;
  } catch (err) {
    console.log(err);
  }
};

const getSignature = () => {
  const nonce = document.getElementById("nonce").value;
  const publicAddress = document.getElementById("address").value;
  web3.eth.personal.sign(
    web3.utils.utf8ToHex(nonce),
    publicAddress,
    (err, signature) => {
      if (err) {
        console.log(err);
      } else {
        document.getElementById("sign").value = signature;
      }
    }
  );
};

const verifySignature = async () => {
  // TO BE IMPLEMENTED
};
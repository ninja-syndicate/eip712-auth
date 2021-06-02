const enableEthereum = async () => {
  if (window.ethereum) {
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
  document.getElementById("address").value = publicAddress;
};

const getMetamaskAddress = async () => {
  // Log all eth accounts
  const accounts = await ethereum.request({ method: "eth_accounts" });

  // Grab the metamask public address
  return accounts[0];
};

const getNonce = async () => {
  try {
    const response = await axios.post("/request_nonce");
    document.getElementById("nonce").value = response.data;
  } catch (err) {
    console.log(err);
  }
};

const getSignature = () => {
  const nonce = document.getElementById("nonce").value;
  const publicAddress = document.getElementById("address").value;
  const companyName = "John";
  const message = `🏆Hi! This is ${companyName}👋!\n\n 🎯Sign this message to prove you have access to this wallet and I’ll log you in. This won’t cost you any Ether.\n
✅To stop others from using your wallet, here’s a unique message ID they can’t guess:\n ${nonce}`;
  web3.eth.personal.sign(
    web3.utils.utf8ToHex(message),
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

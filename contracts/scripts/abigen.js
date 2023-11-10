const fs = require('fs');
const compiledContract = require('../artifacts/contracts/Crater.sol/Crater.json');
const { exec } = require("child_process");

const abi  = compiledContract.abi;
const bin = compiledContract.bytecode.slice(2, compiledContract.bytecode.length);

const basePath = "./artifacts/contracts/Crater.sol/"
fs.writeFileSync(basePath+"crater.abi", JSON.stringify(abi, null, 1));
fs.writeFileSync(basePath+"crater.bin", bin);

exec("abigen --bin "+basePath+"crater.bin --abi "+basePath+"crater.abi --pkg=cratercontract --out=../etherman/cratercontract/cratercontract.go", (error, stdout, stderr) => {
    if (error) {
        console.log(`error: ${error.message}`);
        return;
    }
    if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
    }
    console.log("Go cometcontracts package generated");
});
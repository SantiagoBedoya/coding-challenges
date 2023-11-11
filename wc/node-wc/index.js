#!/usr/bin/env node

const { program, Argument } = require("commander");
const fs = require("fs");
const { processInput } = require("./counter");

program.description("Nodejs wc CLI");

program
  .addArgument(new Argument("[filePath]", "file to read"))
  .option("-c", "count bytes", false)
  .option("-l", "count lines", false)
  .option("-w", "count words", false)
  .option("-m", "count characters", false)
  .action((filePath, options) => {
    let input = "";

    if (filePath) {
      input = fs.readFileSync(filePath, "utf-8");
      processInput(input, { ...options, isFromStdin: false, filePath });
    } else {
      process.stdin.on("readable", () => {
        const chunk = process.stdin.read();
        if (chunk !== null) {
          input += chunk;
        }
      });
      process.stdin.on("end", () => {
        processInput(input, { ...options, isFromStdin: true });
      });
    }
  });

program.parse(process.argv);

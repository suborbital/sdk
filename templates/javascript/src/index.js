import "fastestsmallesttextencoderdecoder-encodeinto/EncoderDecoderTogether.min.js";
import { run } from "./lib";

import { setup, plugin } from "@suborbital/plugin";

const decoder = new TextDecoder();

export function run_e(payload, ident) {
  // Imports will be injected by the runtime
  setup(this.imports, ident);

  const input = decoder.decode(payload);
  const result = run(input);

  plugin.returnResult(result);
}

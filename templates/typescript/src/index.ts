import "fastestsmallesttextencoderdecoder-encodeinto/EncoderDecoderTogether.min.js";
import { run } from "./lib";

import { setup, plugin } from "@suborbital/plugin";

declare global {
  var TextEncoder: any;
  var TextDecoder: any;
}

const decoder = new TextDecoder();

export function run_e(payload: ArrayBuffer, ident: number) {
  // Imports will be injected by the runtime
  // @ts-ignore
  setup(this.imports, ident);

  const input = decoder.decode(payload);
  const result = run(input);

  plugin.returnResult(result);
}

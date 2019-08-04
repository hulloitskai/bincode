import { client } from "@/api";

const bincode = {
  /**
   * @param {string} key The key to use for binary encoding.
   * @param {{ plaintext: string, file: any }} config Configurable options.
   * @returns {Promise<string>} The resulting encoded ciphertext.
   */
  async encode(key, config = {}) {
    const { plaintext, file } = config;
    if (file) throw new Error("bincode: file uploads not yet implemented");

    // Create form.
    const form = new FormData();
    form.set("key", key);
    if (plaintext) {
      const text = plaintext.endsWith("\n") ? plaintext : plaintext + "\n";
      form.set("text", text);
    }

    // Send form.
    const { data } = await client.post("/encode", form, {
      headers: { "Content-Type": "multipart/form-data" },
    });
    return data;
  },

  /**
   * @param {string} ciphertext The binary-form text to decode.
   * @returns {Promise<Blob>} The resulting decoded blob.
   */
  async decode(ciphertext) {
    const { data } = await client.post("/decode", ciphertext, {
      headers: { "Content-Type": "text/plain; charset=utf-8" },
    });
    return data;
  },
};

export default bincode;

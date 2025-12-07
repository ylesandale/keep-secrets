const uint8ArrayToBase64 = (bytes: Uint8Array): string => {
  let binary = ''
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i])
  }
  return btoa(binary)
}

const base64ToUint8Array = (base64: string): Uint8Array => {
  const binary = atob(base64)
  const bytes = new Uint8Array(binary.length)
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i)
  }
  return bytes
}

export const useEncryption = () => {
  const encrypt = async (
    plaintext: string
  ): Promise<{ encrypted: string; key: string }> => {
    const key = await crypto.subtle.generateKey(
      {
        name: 'AES-GCM',
        length: 256,
      },
      true,
      ['encrypt', 'decrypt']
    )

    const iv = crypto.getRandomValues(new Uint8Array(12))

    const encoder = new TextEncoder()
    const data = encoder.encode(plaintext)

    const encrypted = await crypto.subtle.encrypt(
      {
        name: 'AES-GCM',
        iv,
      },
      key,
      data
    )

    const exportedKey = await crypto.subtle.exportKey('raw', key)
    const keyArray = new Uint8Array(exportedKey)

    const combined = new Uint8Array(iv.length + encrypted.byteLength)
    combined.set(iv, 0)
    combined.set(new Uint8Array(encrypted), iv.length)

    const encryptedBase64 = uint8ArrayToBase64(combined)
    const keyBase64 = uint8ArrayToBase64(keyArray)

    return {
      encrypted: encryptedBase64,
      key: keyBase64,
    }
  }

  const decrypt = async (
    encryptedBase64: string,
    keyBase64: string
  ): Promise<string> => {
    try {
      const combined = base64ToUint8Array(encryptedBase64)
      const keyArray = base64ToUint8Array(keyBase64)

      const iv = combined.slice(0, 12)
      const encrypted = combined.slice(12)

      const key = await crypto.subtle.importKey(
        'raw',
        keyArray.buffer as ArrayBuffer,
        {
          name: 'AES-GCM',
          length: 256,
        },
        false,
        ['decrypt']
      )

      const decrypted = await crypto.subtle.decrypt(
        {
          name: 'AES-GCM',
          iv,
        },
        key,
        encrypted
      )

      const decoder = new TextDecoder()
      return decoder.decode(decrypted)
    } catch (error) {
      const message =
        error instanceof Error ? error.message : 'Unknown error'
      throw new Error(`Failed to decrypt: ${message}`)
    }
  }

  return {
    encrypt,
    decrypt,
  }
}

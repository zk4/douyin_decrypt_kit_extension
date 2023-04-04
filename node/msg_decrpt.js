
const crypto = require('crypto');
const base64url = require('base64url');

const RANDOM_BYTES_POS = 32;

class MsgDecrypt {
    constructor(encodingAesKey) {
        if (encodingAesKey.length !== 43) {
            throw new Error('AES key 长度不合法');
        }
        const aesKey = Buffer.from(encodingAesKey + '=', 'base64');

        const keySpec = crypto.createCipheriv('aes-256-cbc', aesKey, aesKey.slice(0, 16));
        const iv = aesKey.slice(0, 16);

        this.cipher = crypto.createDecipheriv('aes-256-cbc', aesKey, iv);
    }

    getLength(orderBytes) {
        const buf = Buffer.from(orderBytes);
        buf.swap32();
        return buf.readInt32BE();
    }

    decode(decrypted) {
        let pad = decrypted[decrypted.length - 1];
        if (pad < 1 || pad > 32) {
            pad = 0;
        }
        return decrypted.slice(0, decrypted.length - pad);
    }

    decrypt(text) {
        let original;
        try {
            const encrypted = base64url.toBuffer(text);
            original = this.cipher.update(encrypted);
            original = Buffer.concat([original, this.cipher.final()]);
        } catch (e) {
            console.error(e);
            throw new Error('解码异常');
        }

        let Content;
        let AppId;
        try {
            const bytes = this.decode(original);

            const pos = bytes.slice(RANDOM_BYTES_POS, RANDOM_BYTES_POS + 4);
            const msgLength = this.getLength(pos);

            Content = bytes.slice(RANDOM_BYTES_POS + 4, RANDOM_BYTES_POS + 4 + msgLength).toString('utf8');
            AppId = bytes.slice(RANDOM_BYTES_POS + 4 + msgLength).toString('utf8');

            console.log(`Content: ${Content}`);
            console.log(`ThirdParty AppID: ${AppId}`);
        } catch (e) {
            console.error(e);
            throw new Error('Buffer异常');
        }
    }
}

const test = new MsgDecrypt('XXX');
test.decrypt('XXX');

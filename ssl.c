#include <openssl/x509.h>
#include <openssl/pem.h>
#include <openssl/bio.h>
#include <string.h>

char* x509_to_pem(X509 *cert) {
    if (cert == NULL) {
        return NULL;
    }

    BIO *bio = BIO_new(BIO_s_mem());  // Create a memory BIO
    if (bio == NULL) {
        return NULL;
    }

    // Write the X509 certificate in PEM format to the BIO
    if (PEM_write_bio_X509(bio, cert) != 1) {
        BIO_free(bio);
        return NULL;
    }

    // Get the data pointer and length from the BIO
    char *pem_data;
    long pem_length = BIO_get_mem_data(bio, &pem_data);

    // Allocate a buffer to hold the PEM string with a null terminator
    char *pem_string = (char *)malloc(pem_length + 1);
    if (pem_string == NULL) {
        BIO_free(bio);
        return NULL;
    }

    // Copy the PEM data into the buffer and null-terminate it
    memcpy(pem_string, pem_data, pem_length);
    pem_string[pem_length] = '\0';

    BIO_free(bio);  // Free the BIO

    return pem_string;  // Return the PEM-encoded string
}

char* x509_to_der(X509 *cert, int *der_length) {
    if (cert == NULL || der_length == NULL) {
        return NULL;
    }

    // Calculate the length of the DER encoding
    int len = i2d_X509(cert, NULL);
    if (len < 0) {
        return NULL;
    }

    // Allocate memory for the DER data
    unsigned char *der_data = (unsigned char *)malloc(len);
    if (der_data == NULL) {
        return NULL;
    }

    // Convert the X509 certificate to DER format
    unsigned char *p = der_data;
    len = i2d_X509(cert, &p);
    if (len < 0) {
        free(der_data);
        return NULL;
    }

    *der_length = len;  // Set the length of DER data
    return der_data;    // Return DER-encoded data
}


char* convert_x509(X509 *cert, int *der_length) {
	char* tmp = x509_to_der(cert, der_length);
	X509_free(cert);
	return tmp;
}

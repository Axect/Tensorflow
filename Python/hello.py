import tensorflow as tf


def main():
    hello = tf.constant("Hello, TensorFlow!")
    sess = tf.Session()
    
    # In Python3, should use str(~~, encoding="utf-8")
    print(str(sess.run(hello), encoding="utf-8"))

if __name__ == "__main__":
    main()

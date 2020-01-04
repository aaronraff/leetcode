class Foo {
    private Semaphore s2;
    private Semaphore s3;
    
    public Foo() {
        this.s2 = new Semaphore(0, true);
        this.s3 = new Semaphore(0, true);
    }

    public void first(Runnable printFirst) throws InterruptedException {
        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        s2.release();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        // printSecond.run() outputs "second". Do not change or remove this line.
        s2.acquire();
        printSecond.run();
        s3.release();
    }

    public void third(Runnable printThird) throws InterruptedException {
        // printThird.run() outputs "third". Do not change or remove this line.
        s3.acquire();
        printThird.run();
    }
}

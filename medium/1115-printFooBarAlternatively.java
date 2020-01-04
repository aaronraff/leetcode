class FooBar {
    private int n;
    private Semaphore s1;
    private Semaphore s2;
    
    public FooBar(int n) {
        this.n = n;
        this.s1 = new Semaphore(1, true);
        this.s2 = new Semaphore(0, true);
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            // printFoo.run() outputs "foo". Do not change or remove this line.
            this.s1.acquire();
            printFoo.run();
            this.s2.release();
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            // printBar.run() outputs "bar". Do not change or remove this line.
            this.s2.acquire();
            printBar.run();
            this.s1.release();
        }
    }
}
